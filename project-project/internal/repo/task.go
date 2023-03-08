package repo

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int) ([]mtask.VexTaskStagesTemplate, error)
}
