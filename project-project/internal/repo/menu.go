package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data"
)

type MenuRepo interface {
	// FindMenus 查询菜单
	FindMenus(ctx context.Context) ([]*data.ProjectMenu, error)
}
