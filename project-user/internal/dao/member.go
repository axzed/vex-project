package dao

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/gorm"
	"github.com/axzed/project-user/internal/database/interface/conn"
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
func (m MemberDao) SaveMember(conn conn.DbConn, ctx context.Context, mem *data.Member) error {
	// conn.(*gorm.GormConn) 外部事务conn转换为gorm conn
	m.conn = conn.(*gorm.GormConn)
	return m.conn.Tx(ctx).Create(mem).Error
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
