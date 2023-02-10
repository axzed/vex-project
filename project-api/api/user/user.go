package user

import (
	"context"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-user/pkg/dao"
	"github.com/axzed/project-user/pkg/repo"
	login_service_v1 "github.com/axzed/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HandlerUser Handler接口实现类
type HandlerUser struct {
	cache repo.Cache // 缓存
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc, // 缓存(给repo.Cache接口一个具体的dao.Rc实现) 要替换只需要换这里的接口实现
	}
}

// getCaptcha 获取验证码
func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	captchaResponse, err := LoginServiceClient.GetCaptcha(c, &login_service_v1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(captchaResponse.Code))
}
