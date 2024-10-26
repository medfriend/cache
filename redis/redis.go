package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

var redisClient *redis.Client

func GetRedisClient(consulClient *api.Client) *redis.Client {

	cache, _ := consul.GetKeyValue(consulClient, "CACHE")

	var result map[string]string

	err := json.Unmarshal([]byte(cache), &result)

	if err != nil {
		log.Fatalf("Error converting JSON string to map: %v", err)
	}

	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%s",
				result["REDIS_HOST"],
				result["REDIS_PORT"],
			),
			Password: "",
			DB:       0,
		})
	}

	return redisClient
}

func NewCacheProxy(client *api.Client) *CacheProxy {
	return &CacheProxy{client: GetRedisClient(client)}
}

func (p *CacheProxy) GetData(key string) (string, error) {

	val, err := p.client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
