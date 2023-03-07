package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/project"
	"github.com/axzed/project-project/internal/database/gorm"
)

type ProjectDao struct {
	conn *gorm.GormConn
}

// FindProjectByMemId 查询项目 分页
func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProAndMember, int64, error) {
	var pms []*project.ProAndMember
	session := p.conn.Session(ctx)
	db := session.Raw("select * from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? limit ?, ?", memId, (page-1)*size, size)
	err := db.Scan(&pms).Error
	var total int64
	session.Model(&project.ProAndMember{}).Where("member_code = ?", memId).Count(&total)
	return pms, total, err
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorm.NewGormConn(),
	}
}
