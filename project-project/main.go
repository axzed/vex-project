package main

import (
	common "github.com/axzed/project-common"
	_ "github.com/axzed/project-project/api"
	"github.com/axzed/project-project/config"
	"github.com/axzed/project-project/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由初始化
	router.InitRouter(r)
	// 初始化rpc调用
	router.InitUserRpc()
	// grpc初始化
	grpc := router.RegisterGrpc()
	// grpc服务注册到etcd
	router.RegisterEtcdServer()
	stop := func() {
		grpc.Stop()
	}
	// 将优雅启停抽取到common的Run中
	common.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr, stop)
}
