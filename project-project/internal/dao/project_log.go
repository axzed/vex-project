package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorms"
)

type ProjectLogDao struct {
	conn *gorms.GormConn
}

// FindLogByMemberCode 根据用户id查询日志 -> 用于首页动态列表展示项目log
func (p *ProjectLogDao) FindLogByMemberCode(ctx context.Context, memberId int64, page int64, size int64) (list []*data.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	offset := (page - 1) * size
	err = session.Model(&data.ProjectLog{}).
		Where("member_code=?", memberId).
		Limit(int(size)).
		Offset(int(offset)).Order("create_time desc").Find(&list).Error
	err = session.Model(&data.ProjectLog{}).
		Where("member_code=?", memberId).Count(&total).Error
	return
}

// SaveProjectLog 保存项目日志
func (p *ProjectLogDao) SaveProjectLog(pl *data.ProjectLog) {
	session := p.conn.Session(context.Background())
	session.Save(&pl)
}

// FindLogByTaskCode 根据任务id查询日志 -> 不分页
func (p *ProjectLogDao) FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) (list []*data.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	model := session.Model(&data.ProjectLog{})
	if comment == 1 {
		model.Where("source_code=? and is_comment=?", taskCode, comment).Find(&list)
		model.Where("source_code=? and is_comment=?", taskCode, comment).Count(&total)
	} else {
		model.Where("source_code=?", taskCode).Find(&list)
		model.Where("source_code=?", taskCode).Count(&total)
	}
	return
}

// FindLogByTaskCodePage 根据任务id查询日志 -> 分页
func (p *ProjectLogDao) FindLogByTaskCodePage(ctx context.Context, taskCode int64, comment int, page int, pageSize int) (list []*data.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	model := session.Model(&data.ProjectLog{})
	offset := (page - 1) * pageSize
	if comment == 1 {
		model.Where("source_code=? and is_comment=?", taskCode, comment).Limit(pageSize).Offset(offset).Find(&list)
		model.Where("source_code=? and is_comment=?", taskCode, comment).Count(&total)
	} else {
		model.Where("source_code=?", taskCode).Limit(pageSize).Offset(offset).Find(&list)
		model.Where("source_code=?", taskCode).Count(&total)
	}
	return
}

func NewProjectLogDao() *ProjectLogDao {
	return &ProjectLogDao{
		conn: gorms.NewGormConn(),
	}
}
