package models

// Order represents an order record in the database
type Order struct {
    ID         uint    `json:"id" gorm:"primaryKey"`
    UserID     uint    `json:"user_id"`            // References the user making the order
    BookID     uint    `json:"book_id"`            // References the book being ordered
    Quantity   int     `json:"quantity"`           // Number of units ordered
    TotalPrice float64 `json:"total_price"`        // Calculated price for the order
    Status     string  `json:"status"`             // Order status: "pending", "completed", "cancelled"
}

// CreateOrderInput is used for incoming order creation requests
type CreateOrderInput struct {
    BookID   uint `json:"book_id" binding:"required"`
    Quantity int  `json:"quantity" binding:"required"`
}

