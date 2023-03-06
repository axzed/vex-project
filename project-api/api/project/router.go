package project

import (
	"github.com/axzed/project-api/middleware"
	"github.com/axzed/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init project router")
	// 注册路由 将当前路由接口实现类append进routers切片
	router.Register(&RouterProject{})
}

// RouterProject Router路由接口实现类
type RouterProject struct {
}

// Route implement Router interface
func (*RouterProject) Route(r *gin.Engine) {
	// 初始化grpc的客户端连接
	InitProjectRpcClient()
	h := NewHandlerProject()
	group := r.Group("/project/index")
	// 接口定义处
	// 路由注册
	// TokenVerify()中间件 用于验证token
	group.Use(middleware.TokenVerify())
	group.POST("", h.index)
}
