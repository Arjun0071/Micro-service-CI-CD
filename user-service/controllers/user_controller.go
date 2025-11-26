package controllers

import (
	"net/http"
	"user-service/models"
	"user-service/utils"
        "log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("/data/users.db"), &gorm.Config{})
        if err != nil {
            log.Fatalf("Failed to connect database: %v", err) // Detailed error
        }
	// AutoMigrate separately and check for errors
	err = db.AutoMigrate(&models.User{})
        if err != nil {
            log.Fatalf("Failed to migrate user model: %v", err)
        }
}

// REGISTER USER
func RegisterUser(c *gin.Context) {
    var input models.RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hash, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user := models.User{
        Username:     input.Username,
        Email:        input.Email,
        PasswordHash: hash,
    }

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// LOGIN
func Login(c *gin.Context) {
    var input models.LoginInput  // <-- instead of inline struct

    // Bind JSON into LoginInput struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    result := db.Where("email = ?", input.Email).First(&user)

    // Check DB result AND password match
    if result.Error != nil || !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateToken(user.ID)

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// GET USER (Protected)
func GetUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User
	db.First(&user, userID)

	c.JSON(http.StatusOK, user)
}

// UPDATE USER (Protected)
func UpdateUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User
	db.First(&user, userID)

	var input models.User
	c.ShouldBindJSON(&input)

	user.Username = input.Username
	user.Email = input.Email

	db.Save(&user)

	c.JSON(http.StatusOK, user)
}

