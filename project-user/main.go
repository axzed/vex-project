package main

import (
	"github.com/axzed/project-common"
	"github.com/axzed/project-common/logs"
	"github.com/axzed/project-user/config"
	"github.com/axzed/project-user/router"
	"github.com/gin-gonic/gin"
	"log"

	_ "github.com/axzed/project-user/api"
)

func main() {
	r := gin.Default()
	// 日志初始化
	lc := &logs.LogConfig{
		DebugFileName: "G:\\vex-project\\back-end\\vex-project\\logs\\project-debug.log",
		InfoFileName:  "G:\\vex-project\\back-end\\vex-project\\logs\\project-info.log",
		WarnFileName:  "G:\\vex-project\\back-end\\vex-project\\logs\\project-warn.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	// 路由初始化
	router.InitRouter(r)
	// 将优雅启停抽取到common的Run中
	common.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr)
}
