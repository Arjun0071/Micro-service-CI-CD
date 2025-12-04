package middlewares

import (
    "book-service/metrics"
    "time"
    "fmt"

    "github.com/gin-gonic/gin"
)

func MetricsMiddleware(service string) gin.HandlerFunc {
    return func(c *gin.Context) {

        start := time.Now()

        // Increase active connections
        metrics.ActiveConnections.WithLabelValues(service).Inc()

        // Process request
        c.Next()

        // Decrease active connections
        metrics.ActiveConnections.WithLabelValues(service).Dec()

        duration := time.Since(start).Seconds()

        // Labels
        method := c.Request.Method
        path := c.FullPath()
        status := c.Writer.Status()

        // Update metrics
        metrics.RequestCount.WithLabelValues(service, method, path, 
            fmt.Sprintf("%d", status)).Inc()

        metrics.RequestDuration.WithLabelValues(service, method, path).Observe(duration)
    }
}

