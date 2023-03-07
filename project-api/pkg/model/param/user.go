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

// LoginReq 登录请求参数
type LoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

// LoginResp 登录响应参数
type LoginResp struct {
	Member           Member             `json:"member"`
	TokenList        TokenList          `json:"tokenList"`
	OrganizationList []OrganizationList `json:"organizationList"`
}

// Member 用户信息
type Member struct {
	Name             string `json:"name"`
	Mobile           string `json:"mobile"`
	Status           int    `json:"status"`
	Code             string `json:"code"`
	Email            string `json:"email"`
	CreateTime       string `json:"create_time"`
	LastLoginTime    string `json:"last_login_time"`
	OrganizationCode string `json:"organization_code"`
}

// TokenList token信息
type TokenList struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenType      string `json:"tokenType"`
	AccessTokenExp int64  `json:"accessTokenExp"`
}

// OrganizationList 组织信息
type OrganizationList struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	OwnerCode   string `json:"owner_code"`
	CreateTime  string `json:"create_time"`
	Personal    int32  `json:"personal"`
	Address     string `json:"address"`
	Province    int32  `json:"province"`
	City        int32  `json:"city"`
	Area        int32  `json:"area"`
	Code        string `json:"code"`
}
