package user

import (
	common "github.com/axzed/project-common"
	"github.com/axzed/project-user/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	mobile := ctx.PostForm("mobile")
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, resp.Fail(model.ErrNotMobile, "手机号格式错误"))
		return
	}
	// 3. 生成验证码 (随机4位1000-9999或者随机6位100000-999999)
	varifyCode := "123456"
	// 4. 调用短信平台 (三方 放入go协程中执行 不影响主流程 接口可以快速响应)
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("验证码发送成功: ")
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, varifyCode)
	}()
	// 5. 将验证码存入redis (key:手机号 value:验证码 过期时间: 15分钟)
	ctx.JSON(http.StatusOK, resp.Success(varifyCode))
}
