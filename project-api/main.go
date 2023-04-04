package main

import (
	_ "github.com/axzed/project-api/api"
	"github.com/axzed/project-api/api/middleware"
	"github.com/axzed/project-api/config"
	"github.com/axzed/project-api/router"
	common "github.com/axzed/project-common"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 调用接口响应中间件
	r.Use(middleware.RequestLog())
	// 静态文件 上传文件
	r.StaticFS("/upload", http.Dir("upload"))
	// 路由初始化;
	router.InitRouter(r)
	// 开启pprof 默认访问路径 /debug/pprof
	pprof.Register(r)
	// 将优雅启停抽取到common的Run中
	common.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr, nil)
}
