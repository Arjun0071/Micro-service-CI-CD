package controllers

import (
    "log"
    "net/http"
    "time"

    "order-service/models"
    "order-service/utils"
    "order-service/metrics"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes SQLite and performs migration
func InitDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("/data/orders.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
    }
    if err := db.AutoMigrate(&models.Order{}); err != nil {
        log.Fatalf("Failed to migrate order model: %v", err)
    }
}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
    start := time.Now() // For duration metric

    var input models.CreateOrderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        metrics.OrdersFailed.WithLabelValues("order-service").Inc()
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)

    // Check book availability
    available, price, err := utils.BookAvailability(input.BookID, input.Quantity)
    if err != nil || !available {
         metrics.OrdersFailed.WithLabelValues("order-service").Inc()
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order := models.Order{
        UserID:     userID,
        BookID:     input.BookID,
        Quantity:   input.Quantity,
        TotalPrice: float64(input.Quantity) * price,
        Status:     "placed",
    }

    if err := db.Create(&order).Error; err != nil {
        metrics.OrdersFailed.WithLabelValues("order-service").Inc()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    
    metrics.OrdersCreated.WithLabelValues("order-service").Inc()
    metrics.OrdersRevenue.WithLabelValues("order-service").Add(order.TotalPrice)
    metrics.OrderCreationDuration.WithLabelValues("order-service").Observe(time.Since(start).Seconds())

    c.JSON(http.StatusCreated, order)
}

// GetOrder returns order details by ID
func GetOrder(c *gin.Context) {
    id := c.Param("id")
    var order models.Order
    if err := db.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

// GetUserOrders returns all orders for a user
func GetUserOrders(c *gin.Context) {
    userID := c.Param("userId")
    var orders []models.Order
    db.Where("user_id = ?", userID).Find(&orders)
    c.JSON(http.StatusOK, orders)
}

