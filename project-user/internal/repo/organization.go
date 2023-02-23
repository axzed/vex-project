package repo

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/interface/conn"
)

type OrganizationRepo interface {
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error)
	SaveOrganization(conn conn.DbConn, ctx context.Context, org *data.Organization) error
	FindOrganizationByMemberId(ctx context.Context, id int64) ([]*data.Organization, error)
}
