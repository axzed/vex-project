package gorms

import (
	"context"
	"gorm.io/gorm"
)

var _db *gorm.DB

// init 初始化数据库连接
//func init() {
//	//配置MySQL连接参数
//	username := config.AppConf.MysqlConfig.Username //账号
//	password := config.AppConf.MysqlConfig.Password //密码
//	host := config.AppConf.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
//	port := config.AppConf.MysqlConfig.Port         //数据库端口
//	Dbname := config.AppConf.MysqlConfig.Db         //数据库名
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
//	var err error
//	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic("连接数据库失败, error=" + err.Error())
//	}
//}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return _db
}

func SetDB(db *gorm.DB) {
	_db = db
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

// Begin 开启事务
func (g *GormConn) Begin() {
	g.tx = g.db.Begin()
}

// Rollback 回滚事务
func (g *GormConn) Rollback() {
	g.tx.Rollback()
}

// Commit 提交事务
func (g *GormConn) Commit() {
	g.tx.Commit()
}

// Session 获取gorm session
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}

// Tx 获取gorm事务
func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}