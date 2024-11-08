package cacheHandler

import (
	"cache-go/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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

	var jsonData map[string]interface{}

	err = json.Unmarshal([]byte(data), &jsonData)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error al procesar los datos de la cache",
		})
		return
	}

	c.JSON(200, jsonData)
}

func PostCache(c *gin.Context, cacheClient *redis.CacheProxy) {
	ctx := context.Background()

	// Crear la clave para Redis en función de la URL
	pathParts := strings.Split(c.Request.URL.Path, "/")
	service := strings.Join(pathParts[2:], "/")

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &jsonData); err == nil {
		bodyBytes, _ = json.Marshal(jsonData)
	}

	// Almacenar el body en Redis utilizando PostData de CacheProxy
	if err := cacheClient.PostData(ctx, service, string(bodyBytes)); err != nil {
		c.JSON(500, gin.H{"error": "Error al almacenar en Redis"})
		return
	}

	fmt.Println("Datos almacenados en cache con la clave:", service)
	c.JSON(200, gin.H{"message": "Datos almacenados en cache"})
}
