package config

import (
	"github.com/axzed/project-project/internal/dao"
	"github.com/go-redis/redis/v8"
)

// ReConnRedis 重连redis
func (c *Config) ReConnRedis() {
	rdb := redis.NewClient(c.InitRedisOptions())
	rc := &dao.RedisCache{
		Rdb: rdb,
	}
	dao.Rc = rc
}
