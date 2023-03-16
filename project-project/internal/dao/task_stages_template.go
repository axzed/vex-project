package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/gorm"
)

type TaskStagesTemplateDao struct {
	conn *gorm.GormConn
}

// FindByProjectTemplateId 根据项目模板id查找任务阶段模板
func (t *TaskStagesTemplateDao) FindByProjectTemplateId(ctx context.Context, projectTemplateCode int) (list []*mtask.VexTaskStagesTemplate, err error) {
	session := t.conn.Session(ctx)
	err = session.
		Model(&mtask.VexTaskStagesTemplate{}).
		Where("project_template_code = ?", projectTemplateCode).
		Order("sort desc, id asc").
		Find(&list).
		Error
	return list, err

}

// FindInProTemIds 查找项目模板下的所有任务阶段模板
func (t *TaskStagesTemplateDao) FindInProTemIds(ctx context.Context, ids []int) ([]mtask.VexTaskStagesTemplate, error) {
	var tsts []mtask.VexTaskStagesTemplate
	session := t.conn.Session(ctx)
	err := session.
		Model(&mtask.VexTaskStagesTemplate{}).
		Where("project_template_code in ?", ids).
		Find(&tsts).
		Error
	return tsts, err
}

func NewTaskStagesTemplateDao() *TaskStagesTemplateDao {
	return &TaskStagesTemplateDao{
		conn: gorm.NewGormConn(),
	}
}
