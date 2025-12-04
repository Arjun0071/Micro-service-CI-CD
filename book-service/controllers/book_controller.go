package controllers

import (
    "book-service/models"
    "book-service/metrics"
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

// Initialize DB
func InitDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("/data/books.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
    }

    err = db.AutoMigrate(&models.Book{})
    if err != nil {
        log.Fatalf("Failed to migrate Book model: %v", err)
    }
}

// ---------------------
// Handlers
// ---------------------

func GetBooks(c *gin.Context) {
    var books []models.Book
    db.Find(&books)
    c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    result := db.First(&book, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
    var book models.Book

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Create(&book)

    // ---- Metrics (AFTER DB, BEFORE response) ----
    metrics.BooksCreated.WithLabelValues("book-service").Inc()
    metrics.BookStock.WithLabelValues("book-service", book.Title).Set(float64(book.Stock))
    // ---------------------------------------------

    c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    if err := db.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    var updateData models.Book
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.Title = updateData.Title
    book.Author = updateData.Author
    book.Price = updateData.Price
    book.Stock = updateData.Stock

    db.Save(&book)

    // ---- Metrics ----
    metrics.BooksUpdated.WithLabelValues("book-service").Inc()
    metrics.BookStock.WithLabelValues("book-service", book.Title).Set(float64(book.Stock))
    // ------------------

    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    if err := db.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    db.Delete(&book)

    // ---- Metrics ----
    metrics.BooksDeleted.WithLabelValues("book-service").Inc()
    metrics.BookStock.DeleteLabelValues("book-service", book.Title)
    // ------------------

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

