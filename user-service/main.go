package main

import (
	"user-service/controllers"
	"user-service/routes"
        "user-service/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDB()

	router := gin.Default()

        router.Use(middlewares.CORSMiddleware())

	routes.RegisterRoutes(router)

	router.Run(":8083")
}

