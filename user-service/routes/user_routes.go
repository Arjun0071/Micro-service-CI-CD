package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// Auth routes
	router.POST("/users", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	auth.GET("/users", controllers.GetUser)
	auth.PUT("/users", controllers.UpdateUser)
}

