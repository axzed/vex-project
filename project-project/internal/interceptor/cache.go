package interceptor

import (
	"context"
	"encoding/json"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/repo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type CacheInterceptor struct {
	cache    repo.Cache     // 缓存实例
	cacheMap map[string]any // 缓存map (key为请求参数, value为返回值) 对应rpc请求参数和返回值
}

type CacheRespOption struct {
	path   string
	typ    any
	expire time.Duration
}

func New() *CacheInterceptor {
	// 初始化缓存map
	cacheMap := make(map[string]any)
	//cacheMap["/project.ProjectService/FindProjectByMemId"] = &project.MyProjectResponse{}
	return &CacheInterceptor{
		cache:    dao.Rc,
		cacheMap: cacheMap,
	}
}

// Cache 缓存拦截器
// grpc.ServerOption 为grpc的ServerOption类型 (grpc.ServerOption是一个函数类型)
func (c *CacheInterceptor) Cache() grpc.ServerOption {
	// grpc.UnaryInterceptor 为grpc的UnaryInterceptor类型 (grpc.UnaryInterceptor是一个函数类型)
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		respType := c.cacheMap[info.FullMethod]
		// 该接口请求不需要缓存 直接返回
		if respType == nil {
			return handler(ctx, req)
		}
		// 先查询是否有缓存 有的话 直接返回 没有的话 放入缓存
		con, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		marshal, _ := json.Marshal(req)
		cacheKey := encrypts.Md5(string(marshal))
		// 从缓存中获取
		respJson, _ := c.cache.Get(con, info.FullMethod+"::"+cacheKey)
		// 如果有缓存 直接返回
		if respJson != "" {
			json.Unmarshal([]byte(respJson), respType)
			zap.L().Info(info.FullMethod + " 从缓存中获取")
			return respType, nil
		}
		// 缓存没有命中
		resp, err = handler(ctx, req)
		bytes, _ := json.Marshal(resp)
		c.cache.Put(con, info.FullMethod+"::"+cacheKey, string(bytes), 5*time.Minute)
		zap.L().Info(info.FullMethod + " 放入缓存")
		return
	})
}
