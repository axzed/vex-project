package task_service_v1

import (
	"context"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-common/tms"
	"github.com/axzed/project-grpc/task"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"time"
)

// TaskService 任务服务
type TaskService struct {
	task.UnimplementedTaskServiceServer
	cache                  repo.Cache              // 缓存
	transaction            transaction.Transaction // 事务
	projectRepo            repo.ProjectRepo
	projectTemplateRepo    repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
}

// NewTaskService 初始化服务
func NewTaskService() *TaskService {
	return &TaskService{
		// 为定义的接口赋上实现类
		cache:                  dao.Rc,
		transaction:            dao.NewTransactionImpl(),
		projectRepo:            dao.NewProjectDao(),
		projectTemplateRepo:    dao.NewProjectTemplateDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
		taskStagesRepo:         dao.NewTaskStagesDao(),
	}
}

// TaskStages 任务阶段
func (t *TaskService) TaskStages(ctx context.Context, msg *task.TaskReqMessage) (*task.TaskStagesResponse, error) {
	projectCode := encrypts.DecryptNoErr(msg.ProjectCode)
	page := msg.Page
	pageSize := msg.PageSize

	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stages, total, err := t.taskStagesRepo.FindStagesByProjectId(c, projectCode, page, pageSize)
	if err != nil {
		zap.L().Error("taskStagesRepo.FindStagesByProjectId", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}

	var tsMessages []*task.TaskStagesMessage
	copier.Copy(&tsMessages, stages)
	if tsMessages == nil {
		return &task.TaskStagesResponse{List: tsMessages, Total: 0}, nil
	}
	stagesMap := mtask.ToTaskStagesMap(stages)
	// 循环赋值
	for _, v := range tsMessages {
		taskStages := stagesMap[int(v.Id)]
		v.Code = encrypts.EncryptNoErr(int64(v.Id))
		v.CreateTime = tms.FormatByMill(taskStages.CreateTime)
		v.ProjectCode = msg.ProjectCode
	}

	return &task.TaskStagesResponse{List: tsMessages, Total: total}, nil
}
