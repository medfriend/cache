package httpServer

import (
	"cache-go/httpServer/cacheHandler"
	"cache-go/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitHttpServer(cacheClient *redis.CacheProxy, taskQueue chan *http.Request) {

	path := fmt.Sprintf("/%s/*path", os.Getenv("SERVICE_PATH"))

	fmt.Println(path)

	r := gin.Default()

	r.Any(path, func(c *gin.Context) {

		c.Header("Content-Type", "application/json")
		taskQueue <- c.Request

		switch c.Request.Method {
		case "GET":
			cacheHandler.GetCache(c, cacheClient)
		}
	})

	err := r.Run(":8090")

	if err != nil {
		return
	}
}
