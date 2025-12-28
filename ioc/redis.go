package ioc

import (
	"github.com/jwcen/mars/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	opts := &redis.Options{
		Addr: config.Config.Redis.Addr,
		
	}
	if config.Config.Redis.Password != "" {
		opts.Password = config.Config.Redis.Password
	}
	redisClient := redis.NewClient(opts)
	return redisClient
}
