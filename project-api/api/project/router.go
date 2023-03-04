package project

import (
	"github.com/axzed/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init project router")
	// 注册路由 将当前路由接口实现类append进routers切片
	router.Register(&RouterProject{})
}

// RouterUser Router路由接口实现类
type RouterProject struct {
}

// implement Router interface
func (*RouterProject) Route(r *gin.Engine) {
	// 初始化grpc的客户端连接
	InitUserRpcClient()
	h := NewHandlerProject()
	// 接口定义处
	// 路由注册
	r.POST("/project/index", h.index)
}
