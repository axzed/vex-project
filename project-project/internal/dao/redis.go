package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//func init() {
//	// 初始化redis连接
//	rdb := redis.NewClient(config.AppConf.InitRedisOptions())
//	// 初始化redis缓存实例Rc
//	Rc = &RedisCache{Rdb: rdb}
//}

// RedisCache 的一个实例
var Rc *RedisCache

type RedisCache struct {
	Rdb *redis.Client // redis连接实例rdb
}

// Put 将key-value存入redis (设置过期时间)
func (r *RedisCache) Put(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.Rdb.Set(ctx, key, value, expiration).Err()
	return err
}

// Get 从redis中获取key对应的value
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := r.Rdb.Get(ctx, key).Result()
	return result, err
}
