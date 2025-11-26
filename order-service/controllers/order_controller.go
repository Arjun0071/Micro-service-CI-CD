package controllers

import (
    "net/http"
    "order-service/models"
    "order-service/utils"
    "log"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

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
    var input models.CreateOrderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)

    // Check book availability
    available, price, err := utils.BookAvailability(input.BookID, input.Quantity)
    if err != nil || !available {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order := models.Order{
        UserID:     userID,
        BookID:     input.BookID,
        Quantity:   input.Quantity,
        TotalPrice: float64(input.Quantity) * price,

// Note: In a real system, order would remain "pending" until payment confirmation.
// For this simplified version, we treat a successful availability check as "order placed".

        Status:     "placed",
    }

    if err := db.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

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

