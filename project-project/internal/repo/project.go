package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/mproject"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, id int64, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
}
