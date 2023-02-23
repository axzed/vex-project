package repo

import (
	"context"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/interface/conn"
)

// MemberRepo 会员仓库接口
type MemberRepo interface {
	// GetMemberByEmail 根据邮箱获取会员
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	// GetMemberByAccount 根据手机号获取会员
	GetMemberByAccount(ctx context.Context, name string) (bool, error)
	// GetMemberByMobile 根据手机号获取会员
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	// SaveMember 保存会员
	SaveMember(conn conn.DbConn, ctx context.Context, mem *data.Member) error
	FindMember(ctx context.Context, account string, pwd string) (mem *data.Member, err error)
}
