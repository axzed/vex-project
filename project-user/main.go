package main

import (
	"github.com/axzed/project-common"
	"github.com/axzed/project-user/router"
	"github.com/gin-gonic/gin"

	_ "github.com/axzed/project-user/api"
)

func main() {
	r := gin.Default()
	// 路由初始化
	router.InitRouter(r)
	// 将优雅启停抽取到common的Run中
	common.Run(r, "project-user", ":80")
}
