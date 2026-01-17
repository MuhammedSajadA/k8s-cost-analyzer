package main

import (
	"fmt"

	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	// fmt.Printf("test: %s", config.GetEnv("PORT", "8080"))
	config.ConnectDb()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(fmt.Sprintf(":%s", config.GetEnv("PORT", "8080")))
}
