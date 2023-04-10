package domain

import (
	"context"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{
		taskRepo: dao.NewTaskDao(),
	}
}

// FindProjectIdByTaskId 根据任务id获取项目id
func (d *TaskDomain) FindProjectIdByTaskId(taskId int64) (int64, bool, *errs.BError) {
	task, err := d.taskRepo.FindTaskById(context.Background(), taskId)
	if err != nil {
		return 0, false, model.ErrDBFail
	}
	if task == nil {
		return 0, false, nil
	}
	return task.ProjectCode, true, nil
}
