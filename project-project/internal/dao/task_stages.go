package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type TaskStagesDao struct {
	conn *gorm.GormConn
}

// SaveTaskStages 保存任务阶段
func (t *TaskStagesDao) SaveTaskStages(ctx context.Context, conn conn.DbConn, ts *mtask.TaskStages) error {
	// 事务经典操作
	t.conn = conn.(*gorm.GormConn)
	err := t.conn.Tx(ctx).Save(&ts).Error
	return err
}

func NewTaskStagesDao() *TaskStagesDao {
	return &TaskStagesDao{
		conn: gorm.NewGormConn(),
	}
}
