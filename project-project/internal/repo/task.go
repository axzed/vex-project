package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int) ([]mtask.VexTaskStagesTemplate, error)
	FindByProjectTemplateId(ctx context.Context, projectTemplateCode int) (list []*mtask.VexTaskStagesTemplate, err error)
}

type TaskStagesRepo interface {
	SaveTaskStages(ctx context.Context, conn conn.DbConn, ts *mtask.TaskStages) error
	FindStagesByProjectId(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*mtask.TaskStages, total int64, err error)
	FindById(ctx context.Context, id int) (ts *mtask.TaskStages, err error)
}

type TaskRepo interface {
	FindTaskByStageCode(ctx context.Context, stageCode int) (list []*mtask.Task, err error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *mtask.TaskMember, err error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (v *int, err error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (v *int, err error)
	SaveTask(ctx context.Context, conn conn.DbConn, ts *mtask.Task) error
	SaveTaskMember(ctx context.Context, conn conn.DbConn, tm *mtask.TaskMember) error
	FindTaskById(ctx context.Context, taskCode int64) (ts *mtask.Task, err error)
	UpdateTaskSort(ctx context.Context, conn conn.DbConn, ts *mtask.Task) error
	FindTaskByStageCodeLtSort(ctx context.Context, stageCode int, sort int) (ts *mtask.Task, err error)
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, size int64) ([]*mtask.Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*mtask.Task, total int64, err error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*mtask.Task, total int64, err error)
}
