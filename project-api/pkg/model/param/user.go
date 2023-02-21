package param

import (
	"errors"
	common "github.com/axzed/project-common"
)

// RegisterReq 注册请求参数
type RegisterReq struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

// VerifyPassword 校验密码
func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

// Verify 校验参数
func (r RegisterReq) Verify() error {
	if !common.VerifyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if !r.VerifyPassword() {
		return errors.New("两次密码输入不一致")
	}
	return nil
}
