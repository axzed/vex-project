package dao

import (
	"github.com/axzed/project-user/internal/database/gorm"
	"github.com/axzed/project-user/internal/database/interface/conn"
)

type TransactionImpl struct {
	conn conn.DbConn
}

func (t TransactionImpl) Action(f func(conn conn.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

// NewTransactionImpl 创建事务实例
func NewTransactionImpl() *TransactionImpl {
	return &TransactionImpl{
		conn: gorm.NewTransaction(),
	}
}
