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
	SaveProjectCollect(ctx context.Context, pc *mproject.CollectionProject) error
	DeleteProjectCollect(ctx context.Context, memberId int64, projectCode int64) error
	UpdateProject(ctx context.Context, proj *mproject.Project) error
	FindProjectMemberByPid(ctx context.Context, projectCode int64) (list []*mproject.ProjectMember, total int64, err error)
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]mproject.ProjectTemplate, int64, error)
}
