package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/gorm"
	gorm2 "gorm.io/gorm"
)

type TaskDao struct {
	conn *gorm.GormConn
}

// FindTaskMemberByTaskId 根据任务id查询任务成员
func (t *TaskDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *mtask.TaskMember, err error) {
	err = t.conn.Session(ctx).
		Where("task_code=? and member_code=?", taskCode, memberId).
		Limit(1).
		Find(&task).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// FindTaskByStageCode 根据阶段id查询任务
func (t *TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) (list []*mtask.Task, err error) {
	//select * from ms_task where stage_code=77 and deleted =0 order by sort asc
	session := t.conn.Session(ctx)
	err = session.Model(&mtask.Task{}).
		Where("stage_code=? and deleted =0", stageCode).
		Order("sort asc").
		Find(&list).Error
	return
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		conn: gorm.NewGormConn(),
	}
}
