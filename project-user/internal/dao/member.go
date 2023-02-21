package dao

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/gorm"
)

// MemberDao 成员dao
type MemberDao struct {
	conn *gorm.GormConn
}

// NewMemberDao 创建成员dao实例
func NewMemberDao() *MemberDao {
	return &MemberDao{conn: gorm.NewGormConn()}
}

// SaveMember 保存成员
func (m MemberDao) SaveMember(ctx context.Context, mem *data.Member) error {
	return m.conn.Session(ctx).Create(mem).Error
}

// GetMemberByEmail 根据邮箱获取成员
func (m MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// GetMemberByAccount 根据账号获取成员
func (m MemberDao) GetMemberByAccount(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// GetMemberByMobile 根据手机号获取成员
func (m MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("mobile = ?", mobile).Count(&count).Error
	return count > 0, err
}
