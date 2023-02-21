package repo

import (
	"context"
	"github.com/axzed/project-user/internal/data"
)

type OrganizationRepo interface {
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error)
	SaveOrganization(ctx context.Context, org *data.Organization) error
}
