package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func InitHttpServer() {

	path := fmt.Sprintf("/%s/*path", os.Getenv("SERVICE_PATH"))

	fmt.Println(path)

	r := gin.Default()

	r.Any(path, func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		c.JSON(200, gin.H{
			"test": "Test",
		})
	})

	r.Run(":8090")
}
