package routes

import (
    "order-service/controllers"
    "order-service/middlewares"

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

    // Protected order routes
    auth := router.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    auth.POST("/orders", controllers.CreateOrder)
    auth.GET("/orders/:id", controllers.GetOrder)
    auth.GET("/orders/user/:userId", controllers.GetUserOrders)
}

