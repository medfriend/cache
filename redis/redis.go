package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient *redis.Client

// TODO agregar la variables env
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
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
