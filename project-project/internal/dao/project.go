package dao

import (
	"context"
	"fmt"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorms"
	"github.com/axzed/project-project/internal/database/interface/conn"
	gorm2 "gorm.io/gorm"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

// FindProjectById 通过id查询对应项目
func (p *ProjectDao) FindProjectById(ctx context.Context, projectCode int64) (pj *data.Project, err error) {
	err = p.conn.Session(ctx).Where("id=?", projectCode).Find(&pj).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// FindProjectByIds 通过id查询对应项目 (此处id是pids)
func (p *ProjectDao) FindProjectByIds(ctx context.Context, pids []int64) (list []*data.Project, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.Project{}).Where("id in (?)", pids).Find(&list).Error
	return
}

// FindProjectMemberByPid 根据用户id查询项目
func (p *ProjectDao) FindProjectMemberByPid(ctx context.Context, projectCode int64) (list []*data.ProjectMember, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectMember{}).
		Where("project_code=?", projectCode).
		Find(&list).Error
	err = session.Model(&data.ProjectMember{}).
		Where("project_code=?", projectCode).
		Count(&total).Error
	return
}

// UpdateProject 更新项目具体信息
func (p *ProjectDao) UpdateProject(ctx context.Context, proj *data.Project) error {
	return p.conn.Session(ctx).Updates(&proj).Error
}

// DeleteProjectCollect 删除项目收藏
func (p *ProjectDao) DeleteProjectCollect(ctx context.Context, memberId int64, projectCode int64) error {
	return p.conn.Session(ctx).
		Where("member_code = ? and project_code = ?", memberId, projectCode).
		Delete(&data.CollectionProject{}).
		Error
}

// SaveProjectCollect 保存项目收藏
func (p *ProjectDao) SaveProjectCollect(ctx context.Context, pc *data.CollectionProject) error {
	return p.conn.Session(ctx).Save(&pc).Error
}

// UpdateDeleteProject 更新项目deleted状态 (保证了delete 和 recovery 操作复用)
func (p *ProjectDao) UpdateDeleteProject(ctx context.Context, id int64, deleted bool) error {
	var err error
	session := p.conn.Session(ctx)
	if deleted {
		// 删除
		err = session.Model(&data.Project{}).Where("id = ?", id).Update("deleted", 1).Error
	} else {
		// 恢复
		err = session.Model(&data.Project{}).Where("id = ?", id).Update("deleted", 0).Error
	}
	return err
}

// DeleteProject 删除项目
func (p *ProjectDao) DeleteProject(ctx context.Context, id int64) error {
	err := p.conn.Session(ctx).Model(&data.Project{}).Where("id = ?", id).Update("deleted", 1).Error
	return err
}

// FindCollectByPIdAndMemId 查询项目是否收藏
func (p *ProjectDao) FindCollectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? and b.project_code = ? limit 1")
	raw := session.Raw(sql, projectCode, memberId)
	err := raw.Scan(&count).Error
	return count > 0, err
}

// FindProjectByPIdAndMemId 查询项目
func (p *ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*data.ProAndMember, error) {
	var pm *data.ProAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select a.*, b.project_code, b.member_code, b.join_time, b.is_owner, b.authorize from vex_project a, vex_project_member b where a.id = b.project_code and b.member_code = ? and b.project_code = ? limit 1")
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&pm).Error
	return pm, err
}

// SaveProject 保存项目
func (p *ProjectDao) SaveProject(conn conn.DbConn, ctx context.Context, pr *data.Project) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pr).Error
}

// SaveProjectMember 保存项目成员
func (p *ProjectDao) SaveProjectMember(conn conn.DbConn, ctx context.Context, pm *data.ProjectMember) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pm).Error
}

// FindCollectProjectByMemId 查询收藏项目 分页
func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*data.ProAndMember, int64, error) {
	var pms []*data.ProAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from vex_project where id in (select project_code from vex_project_collection where member_code = ?) order by sort limit ?, ?")
	db := session.Raw(sql, memberId, (page-1)*size, size)
	err := db.Scan(&pms).Error
	var total int64
	query := fmt.Sprintf("member_code = ?")
	session.Model(&data.CollectionProject{}).Where(query, memberId).Count(&total)
	return pms, total, err
}

// FindProjectByMemId 查询项目 分页
func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*data.ProAndMember, int64, error) {
	var pms []*data.ProAndMember
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
		conn: gorms.NewGormConn(),
	}
}
