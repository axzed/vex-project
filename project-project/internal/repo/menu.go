package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/menu"
)

type MenuRepo interface {
	// FindMenus 查询菜单
	FindMenus(ctx context.Context) ([]*menu.ProjectMenu, error)
}
