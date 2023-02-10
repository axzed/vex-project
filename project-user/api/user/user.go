package user

import (
	common "github.com/axzed/project-common"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Result{}
	ctx.JSON(200, resp.Success("123456"))
}
