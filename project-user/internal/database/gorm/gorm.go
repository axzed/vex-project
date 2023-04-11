package gorm

import (
	"context"
	"fmt"
	"github.com/axzed/project-project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

// TODO 这个init使用gorm完成了读写分离 但数据库没有配置主从复制

//func init() {
//	if config.AppConf.DbConfig.Separation {
//		//开启读写分离
//		username := config.AppConf.DbConfig.Master.Username //账号
//		password := config.AppConf.DbConfig.Master.Password //密码
//		host := config.AppConf.DbConfig.Master.Host         //数据库地址，可以是Ip或者域名
//		port := config.AppConf.DbConfig.Master.Port         //数据库端口
//		Dbname := config.AppConf.DbConfig.Master.Db         //数据库名
//		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
//		var err error
//		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//			Logger: logger.Default.LogMode(logger.Info),
//		})
//		if err != nil {
//			panic("连接数据库失败, error=" + err.Error())
//		}
//		//slave
//		replicas := []gorm.Dialector{}
//		for _, v := range config.AppConf.DbConfig.Slave {
//			username := v.Username //账号
//			password := v.Password //密码
//			host := v.Host         //数据库地址，可以是Ip或者域名
//			port := v.Port         //数据库端口
//			Dbname := v.Db         //数据库名
//			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
//			cfg := mysql.Config{
//				DSN: dsn,
//			}
//			replicas = append(replicas, mysql.New(cfg))
//		}
//		_db.Use(dbresolver.Register(dbresolver.Config{
//			Sources: []gorm.Dialector{mysql.New(mysql.Config{
//				DSN: dsn,
//			})},
//			Replicas: replicas,
//			Policy:   dbresolver.RandomPolicy{},
//		}).SetMaxOpenConns(200).SetMaxIdleConns(10))
//	} else {
//		//配置MySQL连接参数
//		username := config.AppConf.MysqlConfig.Username //账号
//		password := config.AppConf.MysqlConfig.Password //密码
//		host := config.AppConf.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
//		port := config.AppConf.MysqlConfig.Port         //数据库端口
//		Dbname := config.AppConf.MysqlConfig.Db         //数据库名
//		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
//		var err error
//		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//			Logger: logger.Default.LogMode(logger.Info),
//		})
//		if err != nil {
//			panic("连接数据库失败, error=" + err.Error())
//		}
//	}
//}

// init 初始化数据库连接
func init() {
	//配置MySQL连接参数
	username := config.AppConf.MysqlConfig.Username //账号
	password := config.AppConf.MysqlConfig.Password //密码
	host := config.AppConf.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
	port := config.AppConf.MysqlConfig.Port         //数据库端口
	Dbname := config.AppConf.MysqlConfig.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return _db
}

// GormConn gorm连接
type GormConn struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewGormConn 创建gorm连接
func NewGormConn() *GormConn {
	return &GormConn{db: GetDB()}
}

// NewTransaction 创建带事务的gorm连接
func NewTransaction() *GormConn {
	return &GormConn{db: GetDB(), tx: GetDB()}
}

func (g *GormConn) Begin() {
	g.tx = g.db.Begin()
}

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}

func (g *GormConn) Commit() {
	g.tx.Commit()
}

// Session 获取gorm session
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
