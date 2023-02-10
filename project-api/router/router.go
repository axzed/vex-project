package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 注册路由实例
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

// implement Router interface
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

// routers 路由接口切片
var routers []Router

// Register 注册路由
// 将不同接口append进routers切片
func Register(ro ...Router) {
	routers = append(routers, ro...)
}

// InitRouter 路由初始
func InitRouter(r *gin.Engine) {
	// 方式一需要在当前函数下注册路由
	//rg := New()
	//// 注册用户模块路由
	//rg.Route(&user.RouterUser{}, r)

	// 方式二可以在每次添加新的路由接口时, 在对应的路由接口文件中注册路由
	// 遍历routers切片, 调用Route方法
	for _, ro := range routers {
		ro.Route(r)
	}
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}
