package main

import "time"

// Bike Contains the static data for a bike
type Bike struct {
	ID           string    `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	ModelYear    int       `json:"modelYear"`
	Price        float32   `json:"price"`
	PurchaseDate time.Time `json:"purchaseDate"`
}
