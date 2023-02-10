package login_service_v1

import (
	"context"
	"errors"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-user/pkg/dao"
	"github.com/axzed/project-user/pkg/repo"
	"go.uber.org/zap"
	"log"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errors.New("手机号格式不正确")
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
		err := ls.cache.Put(ctx, "REGISTER_"+mobile, varifyCode, 15*time.Minute)
		if err != nil {
			log.Printf("验证码放入缓存失败, caused: %v\n", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s\n", mobile, varifyCode)
	}()
	return &CaptchaResponse{Code: varifyCode}, nil
}
