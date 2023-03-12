package login_service_v1

import (
	"context"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-common/jwts"
	"github.com/axzed/project-common/tms"
	"github.com/axzed/project-grpc/user/login"
	"github.com/axzed/project-user/config"
	"github.com/axzed/project-user/internal/dao"
	"github.com/axzed/project-user/internal/data"
	"github.com/axzed/project-user/internal/database/interface/conn"
	"github.com/axzed/project-user/internal/database/interface/transaction"
	"github.com/axzed/project-user/internal/repo"
	"github.com/axzed/project-user/pkg/model"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
	"time"
)

// LoginService 登录服务
type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache              // 缓存
	memberRepo       repo.MemberRepo         // 成员操作
	organizationRepo repo.OrganizationRepo   // 组织操作
	transaction      transaction.Transaction // 事务
}

// NewLoginService LoginService构造函数
func NewLoginService() *LoginService {
	return &LoginService{
		// 为定义的接口赋上实现类
		cache:            dao.Rc,
		memberRepo:       dao.NewMemberDao(),
		organizationRepo: dao.NewOrganizationDao(),
		transaction:      dao.NewTransactionImpl(),
	}
}

// GetCaptcha 获取验证码
func (ls *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.ConvertToGrpcError(model.ErrNotMobile)
	}
	// 3. 生成验证码 (随机4位1000-9999或者随机6位100000-999999)
	varifyCode := "123456"
	// 4. 调用短信平台 (三方 放入go协程中执行 不影响主流程 接口可以快速响应)
	go func() {
		time.Sleep(2 * time.Second)
		// test log
		zap.L().Info("短信平台调用成功，发送短信 INFO")
		// redis 假设后续缓存可能用MySQL, mongo, memcache当中的一种
		// 5. 将验证码存入redis (key:手机号 value:验证码 过期时间: 15分钟)log.Println("验证码发送成功: ")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(ctx, model.RegisterRedisKey+mobile, varifyCode, 15*time.Minute)
		if err != nil {
			log.Printf("验证码放入缓存失败, caused: %v\n", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s\n", mobile, varifyCode)
	}()
	return &login.CaptchaResponse{Code: varifyCode}, nil
}

// Register 注册
func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	c := context.Background()
	// 1.可以再次进行参数校验
	// 2.校验验证码
	value, err := ls.cache.Get(c, model.RegisterRedisKey+msg.Mobile)
	// redis.Nil 代表key不存在
	// 草了，这个bug简直难受
	if err == redis.Nil {
		return nil, errs.ConvertToGrpcError(model.ErrCaptchNotFound)
	}
	if err != nil {
		zap.L().Error("Register redis get error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrRedisFail)
	}
	if value != msg.Captcha {
		return nil, errs.ConvertToGrpcError(model.ErrCaptcha)
	}
	// 3.校验业务逻辑(邮箱是否被注册 账号是否被注册 手机号是否被注册)
	exist, err := ls.memberRepo.GetMemberByEmail(c, msg.Email)
	if err != nil {
		zap.L().Error("Register db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if exist {
		return nil, errs.ConvertToGrpcError(model.ErrEmailExist)
	}

	exist, err = ls.memberRepo.GetMemberByAccount(c, msg.Name)
	if err != nil {
		zap.L().Error("Register db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if exist {
		return nil, errs.ConvertToGrpcError(model.ErrAccountExist)
	}

	exist, err = ls.memberRepo.GetMemberByMobile(c, msg.Mobile)
	if err != nil {
		zap.L().Error("Register db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if exist {
		return nil, errs.ConvertToGrpcError(model.ErrMobileExist)
	}
	// 4.执行业务逻辑 将数据存入member表 生成数据存入organization表
	// 4.1 将数据存入member表
	pwd := encrypts.Md5(msg.Password)
	mem := &data.Member{
		Account:       msg.Name,
		Password:      pwd,
		Name:          msg.Name,
		Mobile:        msg.Mobile,
		Email:         msg.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.Normal,
	}
	// 加入事务控制
	err = ls.transaction.Action(func(conn conn.DbConn) error {
		err = ls.memberRepo.SaveMember(conn, c, mem)
		if err != nil {
			zap.L().Error("Register db save error", zap.Error(err))
			return errs.ConvertToGrpcError(model.ErrDBFail)
		}
		// 4.2 生成数据存入organization表
		org := &data.Organization{
			Name:       mem.Name + "的个人项目",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   model.Personal,
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = ls.organizationRepo.SaveOrganization(conn, c, org)
		if err != nil {
			zap.L().Error("register SaveOrganization db err", zap.Error(err))
			return errs.ConvertToGrpcError(model.ErrDBFail)
		}
		return nil
	})

	// 5.返回结果
	return &login.RegisterResponse{}, nil
}

// Login 登录
func (ls *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	c := context.Background()
	// 1. 去数据库查询账号密码 记得密码要先加密
	pwd := encrypts.Md5(msg.Password)
	mem, err := ls.memberRepo.FindMember(c, msg.Account, pwd)
	if err != nil {
		zap.L().Error("Login db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if mem == nil {
		return nil, errs.ConvertToGrpcError(model.ErrAccountOrPwd)
	}
	memMessage := &login.MemberMessage{}
	err = copier.Copy(memMessage, mem)
	// 将用户ID加密
	memMessage.Code, _ = encrypts.EncryptInt64(mem.Id, model.AESKey)
	// 转换数据类型 mem中的LastLoginTime是int64类型 而memMessage中的LastLoginTime是string类型
	// 通过tms.FormatByMill()方法将int64类型转换为string类型
	// CreateTime也是一样
	memMessage.LastLoginTime = tms.FormatByMill(mem.LastLoginTime)
	memMessage.CreateTime = tms.FormatByMill(mem.CreateTime)
	// 2. 根据用户ID去查询对应的组织
	orgs, err := ls.organizationRepo.FindOrganizationByMemberId(c, mem.Id)
	if err != nil {
		zap.L().Error("Login db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	// 将用户ID加密
	for _, v := range orgsMessage {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
		v.OwnerCode = memMessage.Code
		v.CreateTime = tms.FormatByMill(data.ToMap(orgs)[v.Id].CreateTime)
	}
	if len(orgs) > 0 {
		// 获取第一个组织的ID
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	// 3. 用jwt生成token
	memIdStr := strconv.FormatInt(mem.Id, 10)
	exp := time.Duration(config.AppConf.JwtConfig.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.AppConf.JwtConfig.RefreshExp*3600*24) * time.Second
	token := jwts.CreateToken(memIdStr, exp, config.AppConf.JwtConfig.AccessSecret, rExp, config.AppConf.JwtConfig.RefreshSecret)
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessTokenExp: token.AccessExp,
		TokenType:      "bearer",
	}
	// TODO 待优化点: 将 member 和 organization 信息存入redis
	// 4. 返回结果
	return &login.LoginResponse{
		Member:           memMessage,
		OrganizationList: orgsMessage,
		TokenList:        tokenList,
	}, nil
}

// TokenVerify 验证token
func (ls *LoginService) TokenVerify(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	token := msg.Token
	if strings.Contains(token, "bearer") {
		// 去掉bearer
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	parseToken, err := jwts.ParseToken(token, config.AppConf.JwtConfig.AccessSecret)
	if err != nil {
		zap.L().Error("Login TokenVerify ParseToken error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrNotLogin)
	}
	// 从缓存中查询
	// TODO 数据查询 优化点: 登录之后 应该把用户信息缓存起来
	id, _ := strconv.ParseInt(parseToken, 10, 64)
	memberById, err := ls.memberRepo.FindMemberById(context.Background(), id)
	if err != nil {
		zap.L().Error("Login TokenVerify FindMemberById error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	memMessage := &login.MemberMessage{}
	err = copier.Copy(memMessage, memberById)
	// 将用户ID加密
	memMessage.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := ls.organizationRepo.FindOrganizationByMemberId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("Login db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if len(orgs) > 0 {
		// 获取第一个组织的ID
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	memMessage.CreateTime = tms.FormatByMill(memberById.CreateTime)
	return &login.LoginResponse{
		Member: memMessage,
	}, nil
}

// MyOrgList 我的组织列表
func (ls *LoginService) MyOrgList(ctx context.Context, msg *login.UserMessage) (*login.OrgListResponse, error) {
	memId := msg.MemId
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(ctx, memId)
	if err != nil {
		zap.L().Error("MyOrgList FindOrganizationByMemId err", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgsMessage}, nil
}

// FindMemberInfoById 根据用户ID查询用户信息
func (ls *LoginService) FindMemberInfoById(ctx context.Context, msg *login.UserMessage) (*login.MemberMessage, error) {
	// 通过userID查询用户信息
	memberById, err := ls.memberRepo.FindMemberById(context.Background(), msg.MemId)
	if err != nil {
		zap.L().Error("Login TokenVerify FindMemberById error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	memMessage := &login.MemberMessage{}
	err = copier.Copy(memMessage, memberById)
	// 将用户ID加密
	memMessage.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := ls.organizationRepo.FindOrganizationByMemberId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("Login db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if len(orgs) > 0 {
		// 获取第一个组织的ID
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey) // 给用户的组织ID加密
	}
	memMessage.CreateTime = tms.FormatByMill(memberById.CreateTime) // 将用户的创建时间格式化
	return memMessage, nil
}
