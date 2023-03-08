package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/mproject"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, id int64, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
}
