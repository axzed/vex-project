package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/gorm"
)

type TaskStagesTemplateDao struct {
	conn *gorm.GormConn
}

func (t *TaskStagesTemplateDao) FindInProTemIds(ctx context.Context, ids []int) ([]mtask.VexTaskStagesTemplate, error) {
	var tsts []mtask.VexTaskStagesTemplate
	session := t.conn.Session(ctx)
	err := session.Where("project_template_code in ?", ids).Find(&tsts).Error
	return tsts, err
}

func NewTaskStagesTemplateDao() *TaskStagesTemplateDao {
	return &TaskStagesTemplateDao{
		conn: gorm.NewGormConn(),
	}
}