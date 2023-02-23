package config

import (
	"github.com/axzed/project-common/logs"
	"github.com/go-redis/redis/v8"
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
}

// NewConfig 初始化配置
func NewConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// 初始化子配置
	conf.InitServerConfig()
	conf.InitZapLog()
	conf.InitRedisOptions()
	conf.InitGrpcConfig()
	conf.InitEtcdConfig()
	conf.InitMysqlConfig()
	conf.InitJwtConfig()
	// 返回配置好的全局配置
	return conf
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
