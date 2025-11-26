package main

import (
	"book-service/routes"
        "book-service/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
        
        r.Use(middlewares.CORSMiddleware())
	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Metrics (dummy for now)
	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(200, gin.H{"metrics": "not implemented yet"})
	})

	// Register Book Routes
	routes.BookRoutes(r)

	r.Run(":8082")
}

