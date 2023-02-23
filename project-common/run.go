package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 优雅启停函数
// r: gin.Engine实例, srvName: 服务名称, addr: 监听地址
func Run(r *gin.Engine, srvName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// 保证下面的优雅启停
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发 kill -2
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s...", srvName)

	// 优雅关闭服务
	// 5秒内处理完请求
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if stop != nil {
		stop()
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown, caused by: %v", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait time out...")
	}
	log.Printf("%s exiting...", srvName)
}
