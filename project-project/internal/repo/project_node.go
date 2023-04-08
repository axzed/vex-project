package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (list []*data.ProjectNode, err error)
}
