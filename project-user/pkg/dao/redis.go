package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func init() {
	// 初始化redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 初始化redis缓存实例Rc
	Rc = &RedisCache{rdb: rdb}
}

// RedisCache 的一个实例
var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client // redis连接实例rdb
}

// Put 将key-value存入redis (设置过期时间)
func (r *RedisCache) Put(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.rdb.Set(ctx, key, value, expiration).Err()
	return err
}

// Get 从redis中获取key对应的value
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := r.rdb.Get(ctx, key).Result()
	return result, err
}
