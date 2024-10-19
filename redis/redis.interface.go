package redis

import "github.com/redis/go-redis/v9"

type CacheProxy struct {
	client *redis.Client
}
