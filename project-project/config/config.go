package config

import (
	"bytes"
	"github.com/axzed/project-common/logs"
	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"log"
	"os"
)

// AppConf 全局配置实例
var AppConf = NewConfig()

// Config 配置(全局配置)
type Config struct {
	viper       *viper.Viper
	SC          *ServerConfig
	GC          *GrpcConfig
	EtcdConfig  *EtcdConfig
	MysqlConfig *MysqlConfig
	JwtConfig   *JwtConfig
	DbConfig    DbConfig
}

// NewConfig 初始化配置
func NewConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	// 先从 nacos 读取配置 如果读不到 再到本地读
	// 构造 nacos 客户端
	nacosClient := InitNacosClient()
	// 读取 nacos 的配置
	configYaml, err2 := nacosClient.confClient.GetConfig(vo.ConfigParam{
		DataId: "app.yaml",
		Group:  nacosClient.group,
	})
	if err2 != nil {
		log.Fatalln(err2)
	}
	// 实时监听 nacos 的配置
	err2 = nacosClient.confClient.ListenConfig(vo.ConfigParam{
		DataId: "app.yaml",
		Group:  nacosClient.group,
		OnChange: func(namespace, group, dataId, data string) {
			log.Printf("load nacos config changed %s \n", data)
			err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(data)))
			if err != nil {
				log.Printf("load nacos config changed err : %s \n", err.Error())
			}
			//所有的配置应该重新读取
			conf.ReLoadAllConfig()
		},
	})
	if err2 != nil {
		log.Fatalln(err2)
	}
	// 在外面统一设置配置文件格式
	conf.viper.SetConfigType("yaml")
	if configYaml != "" { // 读取 nacos 的配置
		err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(configYaml)))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("load nacos config: %s \n", "success")
	} else { // 读取本地配置
		workDir, _ := os.Getwd()
		conf.viper.SetConfigName("app")
		conf.viper.AddConfigPath(workDir + "/config")
		err := conf.viper.ReadInConfig()
		if err != nil {
			log.Fatalln(err)
		}
	}
	conf.ReLoadAllConfig()
	return conf
}

// ReLoadAllConfig 重新加载所有配置
func (c *Config) ReLoadAllConfig() {
	c.InitServerConfig()
	c.InitZapLog()
	c.InitRedisOptions()
	c.InitGrpcConfig()
	c.InitEtcdConfig()
	c.InitMysqlConfig()
	c.InitJwtConfig()
	c.InitDbConfig()
	// 重新创建数据库连接客户端
	c.ReConnRedis()
	c.ReConnMysql()
}

// ServerConfig 服务配置
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Addr string `mapstructure:"addr"`
}

// GrpcConfig grpc配置
type GrpcConfig struct {
	Name    string `mapstructure:"name"`
	Addr    string `mapstructure:"addr"`
	Version string `mapstructure:"version"`
	Weight  int64  `mapstructure:"weight"`
}

// EtcdConfig etcd配置
type EtcdConfig struct {
	Addrs []string `mapstructure:"addrs"`
}

// MysqlConfig mysql配置
type MysqlConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
}

// DbConfig 读写分离db配置
type DbConfig struct {
	Master     MysqlConfig
	Slave      []MysqlConfig
	Separation bool
}

// JwtConfig jwt配置
type JwtConfig struct {
	AccessExp     int    `mapstructure:"accessExp"`
	RefreshExp    int    `mapstructure:"refreshExp"`
	AccessSecret  string `mapstructure:"accessSecret"`
	RefreshSecret string `mapstructure:"refreshSecret"`
}

// InitJwtConfig 初始化jwt配置
func (c *Config) InitJwtConfig() {
	jc := &JwtConfig{
		AccessExp:     c.viper.GetInt("jwt.accessExp"),
		RefreshExp:    c.viper.GetInt("jwt.refreshExp"),
		AccessSecret:  c.viper.GetString("jwt.accessSecret"),
		RefreshSecret: c.viper.GetString("jwt.refreshSecret"),
	}
	c.JwtConfig = jc
}

// InitServerConfig 初始化服务配置
func (c *Config) InitServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

// InitGrpcConfig 初始化grpc配置
func (c *Config) InitGrpcConfig() {
	gc := &GrpcConfig{}
	gc.Name = c.viper.GetString("grpc.name")
	gc.Addr = c.viper.GetString("grpc.addr")
	gc.Version = c.viper.GetString("grpc.version")
	gc.Weight = c.viper.GetInt64("grpc.weight")
	c.GC = gc
}

// InitZapLog 初始化zap日志
func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

// InitRedisOptions 初始化redis配置
func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("redis.db"),          // use default DB
	}
}

// InitEtcdConfig 初始化etcd配置
func (c *Config) InitEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}

// InitMysqlConfig 初始化mysql配置
func (c *Config) InitMysqlConfig() {
	mc := &MysqlConfig{}
	mc.Username = c.viper.GetString("mysql.username")
	mc.Password = c.viper.GetString("mysql.password")
	mc.Host = c.viper.GetString("mysql.host")
	mc.Port = c.viper.GetInt("mysql.port")
	mc.Db = c.viper.GetString("mysql.db")
	c.MysqlConfig = mc
}

// InitDbConfig 初始化db配置 (配置读写分离)
func (c *Config) InitDbConfig() {
	mc := DbConfig{}
	mc.Separation = c.viper.GetBool("db.separation")
	var slaves []MysqlConfig
	err := c.viper.UnmarshalKey("db.slave", &slaves)
	if err != nil {
		panic(err)
	}
	master := MysqlConfig{
		Username: c.viper.GetString("db.master.username"),
		Password: c.viper.GetString("db.master.password"),
		Host:     c.viper.GetString("db.master.host"),
		Port:     c.viper.GetInt("db.master.port"),
		Db:       c.viper.GetString("db.master.db"),
	}
	mc.Master = master
	mc.Slave = slaves
	c.DbConfig = mc
}
