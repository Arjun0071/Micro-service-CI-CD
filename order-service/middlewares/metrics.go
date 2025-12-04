package middlewares

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
    "order-service/metrics"
)

func MetricsMiddleware(service string) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        metrics.ActiveConnections.WithLabelValues(service).Inc()

        c.Next()

        metrics.ActiveConnections.WithLabelValues(service).Dec()

        duration := time.Since(start).Seconds()
        method := c.Request.Method
        path := c.FullPath()
        if path == "" {
            path = c.Request.URL.Path
        }
        status := c.Writer.Status()

        metrics.RequestCount.WithLabelValues(service, method, path, fmt.Sprintf("%d", status)).Inc()
        metrics.RequestDuration.WithLabelValues(service, method, path).Observe(duration)
    }
}

