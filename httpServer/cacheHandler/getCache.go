package cacheHandler

import (
	"cache-go/redis"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetCache(c *gin.Context, cacheClient *redis.CacheProxy) {

	pathParts := strings.Split(c.Request.URL.Path, "/")
	service := strings.Join(pathParts[2:], "/")

	data, err := cacheClient.GetData(service)

	if err != nil {
		c.JSON(200, gin.H{
			"data": "data no avalible on the cache",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}
