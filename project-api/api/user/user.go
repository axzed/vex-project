package user

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
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
	captchaResponse, err := rpc.LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
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
	_, err = rpc.LoginServiceClient.Register(c, msg)
	if err != nil {
		// 解析grpc中的status错误
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	// 4. 返回响应
	ctx.JSON(http.StatusOK, resp.Success(""))
}

// GetIp 获取客户端ip
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}

// login 登录
func (h *HandlerUser) login(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	var req param.LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数格式有误"))
	}
	// 2. 调用user的grpc 完成登录
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.LoginMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "copy参数失败"))
		return
	}
	msg.Ip = GetIp(ctx)
	loginResp, err := rpc.LoginServiceClient.Login(c, msg)
	if err != nil {
		// 解析grpc中的status错误
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	respData := &param.LoginResp{}
	err = copier.Copy(respData, loginResp)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "copy参数失败"))
		return
	}
	// 3. 返回响应
	ctx.JSON(http.StatusOK, resp.Success(respData))
}

// myOrgList 获取我的组织列表
func (h *HandlerUser) myOrgList(c *gin.Context) {
	result := &common.Result{}
	memberStr, _ := c.Get("memberId")
	memberId := memberStr.(int64)
	list, err2 := rpc.LoginServiceClient.MyOrgList(context.Background(), &login.UserMessage{MemId: memberId})
	if err2 != nil {
		code, msg := errs.ParseGrpcError(err2)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	if list.OrganizationList == nil {
		c.JSON(http.StatusOK, result.Success([]*param.OrganizationList{}))
		return
	}
	var orgs []*param.OrganizationList
	copier.Copy(&orgs, list.OrganizationList)
	c.JSON(http.StatusOK, result.Success(orgs))
}
