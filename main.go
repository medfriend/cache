package main

import (
	"cache-go/consul"
	"cache-go/env"
	"cache-go/httpServer"
	"cache-go/redis"
	"context"
)

var ctx = context.Background()

func main() {
	env.LoadEnv()
	consul.ConnectToConsul()

	httpServer.InitHttpServer(redis.NewCacheProxy())

}
