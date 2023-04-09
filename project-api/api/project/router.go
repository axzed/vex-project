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
	// Auth()中间件 用于验证权限 检验用户是否有权限访问该接口
	group := r.Group("/project")
	group.Use(middleware.TokenVerify())
	group.Use(Auth())
	group.POST("/index", h.index)
	group.POST("/project", h.myProjectList)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.readProject)
	group.POST("/project/recycle", h.recycleProject)
	group.POST("/project/recovery", h.recoveryProject)
	group.POST("/project_collect/collect", h.collectProject)
	group.POST("project/edit", h.editProject)
	group.POST("/project/getLogBySelfProject", h.getLogBySelfProject)
	group.POST("/node", h.nodeList)

	t := NewTask()
	group.POST("/task_stages", t.taskStages)
	group.POST("/project_member/index", t.memberProjectList)
	group.POST("/task_stages/tasks", t.taskList)
	group.POST("/task/save", t.saveTask)
	group.POST("/task/sort", t.taskSort)
	group.POST("/task/selfList", t.myTaskList)
	group.POST("/task/read", t.readTask)
	group.POST("/task_member", t.listTaskMember)
	group.POST("/task/taskLog", t.taskLog)
	group.POST("/task/_taskWorkTimeList", t.taskWorkTimeList)
	group.POST("/task/saveTaskWorkTime", t.saveTaskWorkTime)
	group.POST("/file/uploadFiles", t.uploadFiles)
	group.POST("/task/taskSources", t.taskSources)
	group.POST("/task/createComment", t.createComment)

	a := NewAccount()
	group.POST("/account", a.account)

	d := NewDepartment()
	group.POST("/department", d.department)
	group.POST("/department/save", d.save)
	group.POST("/department/read", d.read)

	auth := NewAuth()
	group.POST("/auth", auth.authList)
	group.POST("/auth/apply", auth.apply)

	menu := NewMenu()
	group.POST("/menu/menu", menu.menuList)
}
