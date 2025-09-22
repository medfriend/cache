package main

import (
	"cache-go/httpServer"
	"cache-go/redis"
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	"github.com/medfriend/shared-commons-go/util/worker"
	"net/http"
	"os"
	"runtime"
)

func main() {
	env.LoadEnv()

	consulIp := os.Getenv("CONSUL_IP")
	consulConn := fmt.Sprint(consulIp, ":8500")

	consulClient := consul.ConnectToConsulKey(consulConn, "CACHE")

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	go httpServer.InitHttpServer(redis.NewCacheProxy(consulClient), taskQueue)

	worker.HandleShutdown(stop, consulClient)
}
