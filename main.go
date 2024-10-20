package main

import (
	"cache-go/consul"
	"cache-go/env"
	"cache-go/httpServer"
	"cache-go/redis"
	"context"
	"fmt"
)

var ctx = context.Background()

func main() {
	env.LoadEnv()
	consul.ConnectToConsul()

	cacheHandler := redis.NewCacheProxy()

	fmt.Println(cacheHandler.GetData("test"))

	httpServer.InitHttpServer()

}
