package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, id int64, page int64, size int64) ([]*mproject.ProAndMember, int64, error)
	SaveProject(conn conn.DbConn, ctx context.Context, pr *mproject.Project) error
	SaveProjectMember(conn conn.DbConn, ctx context.Context, pm *mproject.ProjectMember) error
	FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*mproject.ProAndMember, error)
	FindCollectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error)
	DeleteProject(ctx context.Context, id int64) error
	UpdateDeleteProject(ctx context.Context, id int64, deleted bool) error
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
}
