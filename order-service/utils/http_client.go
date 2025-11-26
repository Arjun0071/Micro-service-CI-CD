package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// BookAvailability checks if a book is available in book-service
func BookAvailability(bookID uint, quantity int) (bool, float64, error) {
    url := fmt.Sprintf("http://book-service:8082/books/%d", bookID)
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
    url := "http://user-service:8083/users"
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
