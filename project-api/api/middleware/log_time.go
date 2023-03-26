package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// RequestLog 接口响应速度中间件
func RequestLog() func(*gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		diff := time.Now().UnixMilli() - start.UnixMilli()
		zap.L().Info(fmt.Sprintf("%s 用时 %d ms", c.Request.RequestURI, diff))
	}
}