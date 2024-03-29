package router

import (
	"github.com/axzed/project-common/discovery"
	"github.com/axzed/project-common/logs"
	"github.com/axzed/project-grpc/account"
	"github.com/axzed/project-grpc/auth"
	"github.com/axzed/project-grpc/department"
	"github.com/axzed/project-grpc/menu"
	"github.com/axzed/project-grpc/project"
	"github.com/axzed/project-grpc/task"
	"github.com/axzed/project-project/config"
	"github.com/axzed/project-project/internal/rpc"
	account_service_v1 "github.com/axzed/project-project/pkg/service/account.service.v1"
	auth_service_v1 "github.com/axzed/project-project/pkg/service/auth.service.v1"
	department_service_v1 "github.com/axzed/project-project/pkg/service/department.service.v1"
	menu_service_v1 "github.com/axzed/project-project/pkg/service/menu.service.v1"
	project_service_v1 "github.com/axzed/project-project/pkg/service/project.service.v1"
	task_service_v1 "github.com/axzed/project-project/pkg/service/task.service.v1"
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
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
	// 遍历routers切片, 调用Route方法
	for _, ro := range routers {
		ro.Route(r)
	}
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

// RegisterGrpc 注册grpc服务
func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.AppConf.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			project.RegisterProjectServiceServer(g, project_service_v1.NewProjectService()) // 注册项目服务
			task.RegisterTaskServiceServer(g, task_service_v1.NewTaskService())             // 注册任务服务
			account.RegisterAccountServiceServer(g, account_service_v1.New())
			department.RegisterDepartmentServiceServer(g, department_service_v1.New())
			auth.RegisterAuthServiceServer(g, auth_service_v1.New())
			menu.RegisterMenuServiceServer(g, menu_service_v1.New())
		}}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			//interceptor.New().CacheInterceptor(),
		)),
	)
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", config.AppConf.GC.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

// RegisterEtcdServer 注册etcd服务
func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	// 注册服务
	// 服务信息
	info := discovery.Server{
		Name:    config.AppConf.GC.Name,
		Addr:    config.AppConf.GC.Addr,
		Version: config.AppConf.GC.Version,
		Weight:  config.AppConf.GC.Weight,
	}
	// 注册服务
	r := discovery.NewRegister(config.AppConf.EtcdConfig.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}

func InitUserRpc() {
	rpc.InitUserRpcClient()
}
