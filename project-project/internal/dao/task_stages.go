package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type TaskStagesDao struct {
	conn *gorm.GormConn
}

// FindById 通过任务步骤id获取任务步骤
func (t *TaskStagesDao) FindById(ctx context.Context, id int) (ts *data.TaskStages, err error) {
	err = t.conn.Session(ctx).Where("id=?", id).Find(&ts).Error
	return
}

// FindStagesByProjectId 根据项目id(projectCode)查询任务阶段
func (t *TaskStagesDao) FindStagesByProjectId(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*data.TaskStages, total int64, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&data.TaskStages{}).
		Where("project_code = ?", projectCode).
		Order("sort asc").
		Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).
		Find(&list).
		Error
	err = session.Model(&data.TaskStages{}).
		Where("project_code = ?", projectCode).
		Count(&total).
		Error
	return
}

// SaveTaskStages 保存任务阶段
func (t *TaskStagesDao) SaveTaskStages(ctx context.Context, conn conn.DbConn, ts *data.TaskStages) error {
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
