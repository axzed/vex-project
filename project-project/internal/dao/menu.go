package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorms"
)

// MenuDao 菜单数据访问对象
type MenuDao struct {
	conn *gorms.GormConn
}

// FindMenus 查询菜单
func (m *MenuDao) FindMenus(ctx context.Context) ([]*data.ProjectMenu, error) {
	var menus []*data.ProjectMenu
	err := m.conn.Session(ctx).Order("pid, sort asc, id asc").Find(&menus).Error
	return menus, err
}

// NewMenuDao 创建菜单数据访问对象
func NewMenuDao() *MenuDao {
	return &MenuDao{
		conn: gorms.NewGormConn(),
	}
}
