package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
)

type ProjectAuthNodeDao struct {
	conn *gorm.GormConn
}

// FindNodeStringList 获取权限节点列表
func (p *ProjectAuthNodeDao) FindNodeStringList(ctx context.Context, authId int64) (list []string, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectAuthNode{}).Where("auth=?", authId).Select("node").Find(&list).Error
	return
}

// DeleteByAuthId 根据权限id删除
func (p *ProjectAuthNodeDao) DeleteByAuthId(ctx context.Context, conn conn.DbConn, authId int64) error {
	p.conn = conn.(*gorm.GormConn)
	tx := p.conn.Tx(ctx)
	err := tx.Where("auth=?", authId).Delete(&data.ProjectAuthNode{}).Error
	return err
}

// Save 保存权限节点
func (p *ProjectAuthNodeDao) Save(ctx context.Context, conn conn.DbConn, authId int64, nodes []string) error {
	p.conn = conn.(*gorm.GormConn)
	tx := p.conn.Tx(ctx)
	var list []*data.ProjectAuthNode
	for _, v := range nodes {
		pn := &data.ProjectAuthNode{
			Auth: authId,
			Node: v,
		}
		list = append(list, pn)
	}
	err := tx.Create(list).Error
	return err
}

func NewProjectAuthNodeDao() *ProjectAuthNodeDao {
	return &ProjectAuthNodeDao{
		conn: gorm.NewGormConn(),
	}
}
