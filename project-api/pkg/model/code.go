package model

import (
	"github.com/axzed/project-common/errs"
)

var (
	ErrNotMobile = errs.NewError(2001, "手机号码格式错误")
)
