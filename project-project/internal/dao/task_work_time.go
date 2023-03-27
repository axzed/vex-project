package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorm"
)

type TaskWorkTimeDao struct {
	conn *gorm.GormConn
}

// Save 保存任务工时
func (t *TaskWorkTimeDao) Save(ctx context.Context, twt *data.TaskWorkTime) error {
	session := t.conn.Session(ctx)
	err := session.Save(&twt).Error
	return err
}

// FindWorkTimeList 查询任务工时列表
func (t *TaskWorkTimeDao) FindWorkTimeList(ctx context.Context, taskCode int64) (list []*data.TaskWorkTime, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&data.TaskWorkTime{}).Where("task_code=?", taskCode).Find(&list).Error
	return
}

func NewTaskWorkTimeDao() *TaskWorkTimeDao {
	return &TaskWorkTimeDao{
		conn: gorm.NewGormConn(),
	}
}
