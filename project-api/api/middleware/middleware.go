package middleware

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetIp 获取ip函数
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}

// TokenVerify Token认证中间件
func TokenVerify() func(*gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		// 1. 从header中获取Token
		token := c.GetHeader("Authorization")
		// 2. 调用user服务进行Token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		// 获取client的ip
		ip := GetIp(c)
		// Token认证
		response, err := rpc.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{
			Token: token,
			Ip:    ip,
		})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		// 3. 认证通过则继续执行 放入gin的上下文，否则返回错误
		c.Set("memberId", response.Member.Id)
		c.Set("memberName", response.Member.Name)
		c.Set("organizationCode", response.Member.OrganizationCode)
		c.Next()
	}
}
