package dao

import (
	"context"
	"fmt"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/database/gorm"
)

type ProjectDao struct {
	conn *gorm.GormConn
}

// FindCollectProjectByMemId 查询收藏项目 分页
func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*mproject.ProAndMember, int64, error) {
	var pms []*mproject.ProAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from vex_project where id in (select project_code from vex_project_collection where member_code = ?) order by sort limit ?, ?")
	db := session.Raw(sql, memberId, (page-1)*size, size)
	err := db.Scan(&pms).Error
	var total int64
	query := fmt.Sprintf("member_code = ?")
	session.Model(&mproject.CollectionProject{}).Where(query, memberId).Count(&total)
	return pms, total, err
}

// FindProjectByMemId 查询项目 分页
func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*mproject.ProAndMember, int64, error) {
	var pms []*mproject.ProAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? %s order by sort limit ?, ?", condition)
	db := session.Raw(sql, memId, (page-1)*size, size)
	err := db.Scan(&pms).Error
	var total int64
	query := fmt.Sprintf("select count(*) from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? %s ", condition)
	tx := session.Raw(query, memId)
	err = tx.Scan(&total).Error
	return pms, total, err
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorm.NewGormConn(),
	}
}
