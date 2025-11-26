package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

// This is used ONLY for receiving signup (registration) JSON input
type RegisterInput struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

// This is used ONLY for receiving login JSON input
type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
