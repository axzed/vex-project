package model

import (
	"github.com/axzed/project-common/errs"
)

var (
	// 业务错误码
	ErrNotMobile    = errs.NewError(10102001, "手机号码格式错误")
	ErrCaptcha      = errs.NewError(10102002, "验证码错误")
	ErrEmailExist   = errs.NewError(10102003, "邮箱已存在")
	ErrAccountExist = errs.NewError(10102004, "账号已存在")
	ErrMobileExist  = errs.NewError(10102005, "手机号已存在")

	// 系统错误码
	ErrRedisFail = errs.NewError(999, "redis操作失败")
	ErrDBFail    = errs.NewError(998, "数据库操作失败")
)
