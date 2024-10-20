package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

var redisClient *redis.Client

func GetRedisClient() *redis.Client {

	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%s",
				os.Getenv("REDIS_HOST"),
				os.Getenv("REDIS_PORT"),
			),
			Password: "",
			DB:       0,
		})
	}

	return redisClient
}

func NewCacheProxy() *CacheProxy {
	return &CacheProxy{client: GetRedisClient()}
}

func (p *CacheProxy) GetData(key string) (string, error) {

	val, err := p.client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
