package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	ctx.JSON(200, "getCaptcha success")
}
