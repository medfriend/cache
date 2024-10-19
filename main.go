package main

import (
	"cache-go/consul"
	"cache-go/env"
	"cache-go/httpServer"
	"cache-go/redis"
	"context"
	"fmt"
	"log"
	"net/http"
)

var ctx = context.Background()

func main() {
	env.LoadEnv()
	consul.ConnectToConsul()

	cacheHandler := redis.NewCacheProxy()

	fmt.Println(cacheHandler.GetData("test"))

	httpServer.InitHttpServer()

	log.Fatal(http.ListenAndServe(":8090", nil))
}
