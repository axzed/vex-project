package main

import (
	common "github.com/axzed/project-common"
	_ "github.com/axzed/project-user/api"
	"github.com/axzed/project-user/config"
	"github.com/axzed/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由初始化
	router.InitRouter(r)
	grpc := router.RegisterGrpc()
	stop := func() {
		grpc.Stop()
	}
	// 将优雅启停抽取到common的Run中
	common.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr, stop)
}
