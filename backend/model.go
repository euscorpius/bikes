package main

import "time"

// User Contains user data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Bike Contains the static data for a bike
type Bike struct {
	ID           string    `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	ModelYear    int       `json:"modelYear"`
	Price        float32   `json:"price"`
	PurchaseDate time.Time `json:"purchaseDate"`
}
