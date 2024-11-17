package redis

import (
	"fmt"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/conf"
	"github.com/go-redis/redis"
)

// NewRedisDB 连接redis
func NewRedisDB(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DB:           0,
		Password:     c.Redis.Password,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}

func GenerateKey(userID int32) string {
	return fmt.Sprintf("auth:%d", userID)
}
