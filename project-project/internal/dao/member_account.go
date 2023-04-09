package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorm"
	gorm2 "gorm.io/gorm"
)

type MemberAccountDao struct {
	conn *gorm.GormConn
}

// FindByMemberId 根据memberId查询
func (m *MemberAccountDao) FindByMemberId(ctx context.Context, memberId int64) (ma *data.MemberAccount, err error) {
	session := m.conn.Session(ctx)
	err = session.Where("member_code=?", memberId).Take(&ma).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// FindList 查询列表
func (m *MemberAccountDao) FindList(ctx context.Context, condition string, organizationCode int64, departmentCode int64, page int64, pageSize int64) (list []*data.MemberAccount, total int64, err error) {
	session := m.conn.Session(ctx)
	offset := (page - 1) * pageSize
	err = session.Model(&data.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Limit(int(pageSize)).Offset(int(offset)).Find(&list).Error
	err = session.Model(&data.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Count(&total).Error
	return
}

func NewMemberAccountDao() *MemberAccountDao {
	return &MemberAccountDao{
		conn: gorm.NewGormConn(),
	}
}
