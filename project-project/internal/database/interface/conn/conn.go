package conn

// DbConn 数据库连接接口
type DbConn interface {
	Begin()
	Rollback()
	Commit()
}
