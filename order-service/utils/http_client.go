package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

// Service URLs loaded from environment variables
var (
    UserServiceURL string
    BookServiceURL string
)

// Init loads environment variables into global variables
func Init() {
    UserServiceURL = os.Getenv("USER_SERVICE_URL")
    BookServiceURL = os.Getenv("BOOK_SERVICE_URL")

    if UserServiceURL == "" {
        fmt.Println("WARNING: USER_SERVICE_URL is not set")
    }
    if BookServiceURL == "" {
        fmt.Println("WARNING: BOOK_SERVICE_URL is not set")
    }
}


// BookAvailability checks if a book is available in book-service
func BookAvailability(bookID uint, quantity int) (bool, float64, error) {
    url := fmt.Sprintf("%s/books/%d", BookServiceURL, bookID)
    resp, err := http.Get(url)
    if err != nil {
        return false, 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return false, 0, fmt.Errorf("book not found or service error")
    }

    var book struct {
        ID    uint    `json:"id"`
        Stock int     `json:"stock"`
        Price float64 `json:"price"`
    }

    body, _ := ioutil.ReadAll(resp.Body)

    if err := json.Unmarshal(body, &book); err != nil {
        return false, 0, err
    }

    if book.Stock < quantity {
        return false, 0, fmt.Errorf("insufficient stock")
    }

    // Everything good → return available = true, price = book.Price
    return true, book.Price, nil
}

// VerifyUserToken calls user-service to validate a JWT token
func VerifyUserToken(token string) (uint, error) {
    url := fmt.Sprintf("%s/users", UserServiceURL)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("invalid or expired token")
    }

    var user struct {
        ID uint `json:"id"`
    }
    body, _ := ioutil.ReadAll(resp.Body)
    if err := json.Unmarshal(body, &user); err != nil {
        return 0, err
    }

    return user.ID, nil
}
