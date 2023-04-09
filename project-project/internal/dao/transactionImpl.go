package dao

import (
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
	"github.com/pkg/errors"
)

type TransactionImpl struct {
	conn conn.DbConn
}

// Action 事务操作 实现事务接口
func (t TransactionImpl) Action(f func(conn conn.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	var bErr *errs.BError
	if errors.Is(err, bErr) {
		bErr = err.(*errs.BError)
		if bErr != nil {
			return bErr
		} else {
			t.conn.Commit()
			return nil
		}
	}
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
