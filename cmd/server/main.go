package main

import (
	"fmt"

	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/config"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/handlers"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	// fmt.Printf("test: %s", config.GetEnv("PORT", "8080"))
	config.ConnectDb()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	router := gin.Default()
	authService := services.NewAuthService(config.DB, config.GetEnv("JWT_SECRET", "secret"))
	authHandler := handlers.NewAuthHandler(authService)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
	router.Run(fmt.Sprintf(":%s", config.GetEnv("PORT", "8080")))
}
