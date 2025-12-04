package main

import (
    "user-service/controllers"
    "user-service/routes"
    "user-service/middlewares"
    "user-service/metrics"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    serviceName := "user-service"

    // Initialize DB
    controllers.InitDB()

    // Initialize Metrics
    metrics.Init(serviceName)

    router := gin.Default()

    // Your existing middlewares
    router.Use(middlewares.CORSMiddleware())

    // Add request metrics middleware
    router.Use(middlewares.MetricsMiddleware(serviceName))

    // Health Check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Prometheus metrics endpoint
    router.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // All application routes
    routes.RegisterRoutes(router)

    router.Run(":8083")
}

