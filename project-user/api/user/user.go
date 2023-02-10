package user

import (
	"context"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-user/pkg/dao"
	"github.com/axzed/project-user/pkg/model"
	"github.com/axzed/project-user/pkg/repo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type HandlerUser struct {
	cache repo.Cache // 缓存
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc, // 缓存(给repo.Cache接口一个具体的dao.Rc实现) 要替换只需要换这里的接口实现
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
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
		// test log
		zap.L().Info("短信平台调用成功，发送短信 INFO")
		zap.L().Debug("短信平台调用成功，发送短信 DEBUG")
		zap.L().Warn("短信平台调用成功，发送短信 WARN")
		// redis 假设后续缓存可能用MySQL, mongo, memcache当中的一种
		// 5. 将验证码存入redis (key:手机号 value:验证码 过期时间: 15分钟)log.Println("验证码发送成功: ")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(ctx, "REGISTER_"+mobile, varifyCode, 15*time.Minute)
		if err != nil {
			log.Printf("验证码放入缓存失败, caused: %v\n", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s\n", mobile, varifyCode)
	}()
	ctx.JSON(http.StatusOK, resp.Success(varifyCode))
}
