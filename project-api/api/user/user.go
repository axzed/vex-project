package user

import (
	"context"
	"github.com/axzed/project-api/pkg/model/param"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/user/login"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

// HandlerUser Handler接口实现类
type HandlerUser struct {
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{}
}

// getCaptcha 获取验证码
func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	captchaResponse, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		// 解析grpc中的status错误
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(captchaResponse.Code))
}

// register 注册
func (h *HandlerUser) register(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	var req param.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数格式有误"))
	}
	// 2. 校验参数
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	// 3. 调用user grpc服务 获取响应
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "copy参数失败"))
	}
	_, err = LoginServiceClient.Register(c, msg)
	if err != nil {
		// 解析grpc中的status错误
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	// 4. 返回响应
	ctx.JSON(http.StatusOK, resp.Success(""))
}
