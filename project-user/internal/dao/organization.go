package dao

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/gorms"
	"github.com/axzed/project-user/internal/database/interface/conn"
)

// OrganizationDao 组织dao
type OrganizationDao struct {
	// 组织dao依赖gorm连接
	conn *gorms.GormConn
}

func (o *OrganizationDao) FindOrganizationByMemberId(ctx context.Context, id int64) ([]*data.Organization, error) {
	var orgs []*data.Organization
	err := o.conn.Session(ctx).Where("member_id=?", id).Find(&orgs).Error
	return orgs, err
}

// NewOrganizationDao 创建组织dao实例
func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.NewGormConn(),
	}
}

// FindOrganizationByMemId 根据成员id获取组织
func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]data.Organization, error) {
	var orgs []data.Organization
	err := o.conn.Session(ctx).Where("member_id=?", memId).Find(&orgs).Error
	return orgs, err
}

// SaveOrganization 保存组织
func (o *OrganizationDao) SaveOrganization(conn conn.DbConn, ctx context.Context, org *data.Organization) error {
	o.conn = conn.(*gorms.GormConn)
	return o.conn.Tx(ctx).Create(org).Error
}
