package startup

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient redis.Cmdable

func InitRedis() redis.Cmdable {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         "localhost:16379",
			DialTimeout:  5 * time.Second, // 连接超时
			ReadTimeout:  3 * time.Second, // 读取超时
			WriteTimeout: 3 * time.Second, // 写入超时
			PoolTimeout:  4 * time.Second, // 连接池超时
		})

		// 使用带超时的 context 进行连接测试
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := redisClient.Ping(ctx).Err()
		if err != nil {
			log.Printf("op=InitRedis||err=%v", err)
			log.Printf("提示: 请确保 Redis 服务正在运行在 localhost:16379")
			panic(err)
		}
		log.Println("Redis 连接成功")
	}
	return redisClient
}
