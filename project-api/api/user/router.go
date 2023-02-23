package user

import (
	"github.com/axzed/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init user router")
	// 注册路由 将当前路由接口实现类append进routers切片
	router.Register(&RouterUser{})
}

// RouterUser Router路由接口实现类
type RouterUser struct {
}

// implement Router interface
func (*RouterUser) Route(r *gin.Engine) {
	// 初始化grpc的客户端连接
	InitUserRpcClient()
	h := NewHandlerUser()
	// 接口定义处
	// 路由注册
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
	r.POST("/project/login", h.login)
}
