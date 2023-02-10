package user

import (
	"github.com/axzed/project-user/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init user router")
	// 注册路由 将当前路由接口实现类append进routers切片
	router.Register(&RouterUser{})
}

// RouterUser Router路由接口实现类
type RouterUser struct {
}

// implement Router interface
func (*RouterUser) Route(r *gin.Engine) {
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
