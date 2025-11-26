package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// Health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Metrics (placeholder)
	router.GET("/metrics", func(c *gin.Context) {
		c.JSON(200, gin.H{"metrics": "coming soon"})
	})

	// Auth routes
	router.POST("/users", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	auth.GET("/users", controllers.GetUser)
	auth.PUT("/users", controllers.UpdateUser)
}

