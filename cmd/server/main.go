package main

import (
	"fmt"

	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/config"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/handlers"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/middleware"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/repositories"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	// fmt.Printf("test: %s", config.GetEnv("PORT", "8080"))
	config.ConnectDb()
}

func main() {
	config.DB.AutoMigrate(&models.User{}, &models.Cluster{})
	router := gin.Default()
	clusterRepo := repositories.NewClusterRepository(config.DB)
	clusterService := services.NewClusterService(clusterRepo)
	clusterHandler := handlers.NewClusterHandler(clusterService)
	namespaceHandler := handlers.NewNamespaceHandler(clusterService)
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
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(config.GetEnv("JWT_SECRET", "secret")))
	{
		protected.POST("/clusters", clusterHandler.AddCluster)
		protected.GET(
			"/clusters/:id/namespaces",
			namespaceHandler.List,
		)
	}

	router.Run(fmt.Sprintf(":%s", config.GetEnv("PORT", "8080")))
}
