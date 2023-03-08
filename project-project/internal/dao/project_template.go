package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/database/gorm"
)

type ProjectTemplateDao struct {
	conn *gorm.GormConn
}

// FindProjectTemplateSystem find system project template
func (p *ProjectTemplateDao) FindProjectTemplateSystem(ctx context.Context, page int64, size int64) (pts []mproject.ProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("is_system = ?", 1).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("is_system = ?", 1).
		Count(&total).
		Error
	return pts, total, err
}

// FindProjectTemplateCustom find custom project template
func (p *ProjectTemplateDao) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) (pts []mproject.ProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("is_system = ? and member_code = ? and organization_code = ?", 0, memId, organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("is_system = ? and member_code = ? and organization_code = ?", 0, memId, organizationCode).
		Count(&total).
		Error
	return pts, total, err
}

// FindProjectTemplateAll find all project template
func (p *ProjectTemplateDao) FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) (pts []mproject.ProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("organization_code = ?", organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&mproject.ProjectTemplate{}).
		Where("organization_code = ?", organizationCode).
		Count(&total).
		Error
	return pts, total, err
}

func NewProjectTemplateDao() *ProjectTemplateDao {
	return &ProjectTemplateDao{
		conn: gorm.NewGormConn(),
	}
}
