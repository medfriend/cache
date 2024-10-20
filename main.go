package main

import (
	"cache-go/httpServer"
	"cache-go/redis"
	"cache-go/util/consul"
	"cache-go/util/env"
	"cache-go/util/worker"
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	env.LoadEnv()
	consulClient := consul.ConnectToConsul()

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	go httpServer.InitHttpServer(redis.NewCacheProxy(), taskQueue)

	worker.HandleShutdown(stop, consulClient)

}
