package dao

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/gorm"
	"github.com/axzed/project-user/internal/database/interface/conn"
	gorm2 "gorm.io/gorm"
)

// MemberDao 成员dao
type MemberDao struct {
	conn *gorm.GormConn
}

func (m *MemberDao) FindMemberByIds(background context.Context, ids []int64) (list []*data.Member, err error) {
	if len(ids) <= 0 {
		return nil, nil
	}
	err = m.conn.Session(background).Model(&data.Member{}).Where("id in (?)", ids).First(&list).Error
	return
}

// FindMemberById 根据id获取会员
func (m *MemberDao) FindMemberById(background context.Context, id int64) (mem *data.Member, err error) {
	err = m.conn.Session(background).Where("id = ?", id).First(&mem).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) FindMember(ctx context.Context, account string, pwd string) (mem *data.Member, err error) {
	err = m.conn.Session(ctx).Where("account = ? and password = ?", account, pwd).First(&mem).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

// NewMemberDao 创建成员dao实例
func NewMemberDao() *MemberDao {
	return &MemberDao{conn: gorm.NewGormConn()}
}

// SaveMember 保存成员
func (m *MemberDao) SaveMember(conn conn.DbConn, ctx context.Context, mem *data.Member) error {
	// conn.(*gorm.GormConn) 外部事务conn转换为gorm conn
	m.conn = conn.(*gorm.GormConn)
	return m.conn.Tx(ctx).Create(mem).Error
}

// GetMemberByEmail 根据邮箱获取成员
func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// GetMemberByAccount 根据账号获取成员
func (m *MemberDao) GetMemberByAccount(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// GetMemberByMobile 根据手机号获取成员
func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&data.Member{}).Where("mobile = ?", mobile).Count(&count).Error
	return count > 0, err
}
