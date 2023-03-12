package dao

import (
	"context"
	"fmt"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type ProjectDao struct {
	conn *gorm.GormConn
}

// FindCollectByPIdAndMemId 查询项目是否收藏
func (p *ProjectDao) FindCollectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select count(*) from vex_project_collection where project_code = ? and member_code = ?")
	raw := session.Raw(sql, projectCode, memberId)
	err := raw.Scan(&count).Error
	return count > 0, err
}

// FindProjectByPIdAndMemId 查询项目
func (p *ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*mproject.ProAndMember, error) {
	var pm *mproject.ProAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? and b.id = ? limit 1")
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&pm).Error
	return pm, err
}

// SaveProject 保存项目
func (p *ProjectDao) SaveProject(conn conn.DbConn, ctx context.Context, pr *mproject.Project) error {
	p.conn = conn.(*gorm.GormConn)
	return p.conn.Tx(ctx).Save(&pr).Error
}

// SaveProjectMember 保存项目成员
func (p *ProjectDao) SaveProjectMember(conn conn.DbConn, ctx context.Context, pm *mproject.ProjectMember) error {
	p.conn = conn.(*gorm.GormConn)
	return p.conn.Tx(ctx).Save(&pm).Error
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
