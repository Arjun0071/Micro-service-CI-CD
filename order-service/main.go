package main

import (
    "order-service/controllers"
    "order-service/routes"
    "order-service/middlewares"
    "order-service/utils"
    "order-service/metrics"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    serviceName := "order-service"

    // Initialize DB
    controllers.InitDB()
    utils.Init()

    metrics.Init(serviceName)

    router := gin.Default()

    // Add metrics middleware
    router.Use(middlewares.MetricsMiddleware(serviceName))

    // CORS middleware if you have it
    router.Use(middlewares.CORSMiddleware())

    // Expose /metrics endpoint
    router.GET("/metrics", gin.WrapH(promhttp.Handler()))
    
    // Health Endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Register application routes
    routes.RegisterRoutes(router)

    // Start service
    router.Run(":8084")
}

