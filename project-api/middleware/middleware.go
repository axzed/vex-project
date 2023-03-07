package middleware

import (
	"context"
	"github.com/axzed/project-api/api/user"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// TokenVerify Token认证中间件
func TokenVerify() func(*gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		// 1. 从header中获取Token
		token := c.GetHeader("Authorization")
		// 2. 调用user服务进行Token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		response, err := user.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{
			Token: token,
		})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		// 3. 认证通过则继续执行 放入gin的上下文，否则返回错误
		c.Set("memberId", response.Member.Id)
		c.Next()
	}
}
