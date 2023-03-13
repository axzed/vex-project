package project

import (
	"github.com/axzed/project-api/api/middleware"
	"github.com/axzed/project-api/api/rpc"
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
	rpc.InitProjectRpcClient()
	h := NewHandlerProject()
	// 路由组
	// 接口定义处
	// 路由注册
	// TokenVerify()中间件 用于验证token
	group := r.Group("/project")
	group.Use(middleware.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project", h.myProjectList)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.readProject)
	group.POST("/project/recycle", h.recycleProject)
	group.POST("/project/recovery", h.recoveryProject)
	group.POST("/project_collect/collect", h.collectProject)
}
