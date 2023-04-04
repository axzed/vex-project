package main

import (
	"fmt"
	_ "github.com/axzed/project-api/api"
	"github.com/axzed/project-api/api/middleware"
	"github.com/axzed/project-api/config"
	"github.com/axzed/project-api/router"
	common "github.com/axzed/project-common"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	//测试代码
	r.GET("/mem", func(c *gin.Context) {
		// 业务代码运行
		outCh := make(chan int)
		// 每秒起10个goroutine，goroutine会阻塞，不释放内存
		tick := time.Tick(time.Second / 10)
		i := 0
		for range tick {
			i++
			fmt.Println(i)
			alloc1(outCh) // 不停的有goruntine因为outCh堵塞，无法释放
		}
	})
	// 将优雅启停抽取到common的Run中
	common.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr, nil)
}

// 一个外层函数
func alloc1(outCh chan<- int) {
	go alloc2(outCh)
}

// 一个内层函数
func alloc2(outCh chan<- int) {
	func() {
		defer fmt.Println("alloc-fm exit")
		// 分配内存，假用一下
		buf := make([]byte, 1024*1024*10)
		_ = len(buf)
		fmt.Println("alloc done")

		outCh <- 0
		//return
	}()
}
