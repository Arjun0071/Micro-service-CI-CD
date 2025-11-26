package main

import (
    "order-service/controllers"
    "order-service/routes"
    "order-service/middlewares"

    "github.com/gin-gonic/gin"
)

func main() {
    controllers.InitDB()

    router := gin.Default()
    router.Use(middlewares.CORSMiddleware())
    routes.RegisterRoutes(router)

    router.Run(":8084") // Port for order-service
}

