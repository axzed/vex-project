package project

import (
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 白名单
var ignores = []string{
	"project/login/register",
	"project/login",
	"project/login/getCaptcha",
	"project/organization",
	"project/auth/apply",
}

// Auth 权限认证
func Auth() func(*gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		// 获取请求的uri
		uri := c.Request.RequestURI
		// 判断是否在忽略列表中
		for _, v := range ignores {
			// 若uri包含忽略列表中的uri，则直接放行
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}

		//判断此uri是否在用户的授权列表中
		a := NewAuth()
		nodes, err := a.GetAuthNodes(c)
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}

		for _, v := range nodes {
			// 若uri包含忽略列表中的uri，则直接放行
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}

		// 若不在忽略列表中，也不在用户的授权列表中，则返回无权限操作
		c.JSON(http.StatusOK, result.Fail(403, "无权限操作"))
		c.Abort()
		return
	}
}
