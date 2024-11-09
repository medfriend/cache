package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

var redisClient *redis.Client

func GetRedisClient(consulClient *api.Client) *redis.Client {

	var cacheKey string

	if os.Getenv("SERVICE_STATUS") == "LOCAL" {
		cacheKey = "REDIS_LOCAL"
	} else {
		cacheKey = "REDIS"
	}

	cache, _ := consul.GetKeyValue(consulClient, cacheKey)

	var result map[string]string

	err := json.Unmarshal([]byte(cache), &result)

	if err != nil {
		log.Fatalf("Error converting JSON string to map: %v", err)
	}

	fmt.Println(result)

	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%s",
				result["REDIS_ADDRESS"],
				result["REDIS_PORT"],
			),
			Password: result["REDIS_PASSWORD"],
			DB:       0,
		})

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("No se pudo conectar a Redis: %v", err)
		} else {
			log.Println("Conexión exitosa a Redis")
		}
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

func (p *CacheProxy) PostData(ctx context.Context, key string, value string) error {
	// Almacenar el valor en Redis
	err := p.client.Set(ctx, key, value, 5*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("error al guardar el dato en Redis: %v", err)
	}

	return nil
}
