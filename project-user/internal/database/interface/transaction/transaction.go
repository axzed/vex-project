package transaction

import (
	"github.com/axzed/project-user/internal/database/interface/conn"
)

// Transaction 事务接口
type Transaction interface {
	Action(func(conn conn.DbConn) error) error
}
