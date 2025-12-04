package main

import (
    "book-service/routes"
    "book-service/middlewares"
    "book-service/metrics"
    "book-service/controllers"
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    service := "book-service"

    // Initialize DB + metrics
    controllers.InitDB()
    metrics.Init(service)

    r := gin.Default()

    // Prometheus HTTP middleware
    r.Use(middlewares.MetricsMiddleware(service))

    // CORS middleware (yours)
    r.Use(middlewares.CORSMiddleware())

    // Health Checks
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Metrics Endpoint
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // Book Routes
    routes.BookRoutes(r)

    r.Run(":8082")
}

