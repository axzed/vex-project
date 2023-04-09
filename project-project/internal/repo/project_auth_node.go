package repo

import (
	"context"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type ProjectAuthNodeRepo interface {
	FindNodeStringList(ctx context.Context, authId int64) ([]string, error)
	DeleteByAuthId(background context.Context, conn conn.DbConn, authId int64) error
	Save(background context.Context, conn conn.DbConn, authId int64, nodes []string) error
}
