package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var bikes []Bike

// GetBikes Get all bikes
func GetBikes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bikes)
}

// GetBike Get a single bike
func GetBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range bikes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Bike{})
}

// CreateBike Add a new bike
func CreateBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bike Bike
	_ = json.NewDecoder(r.Body).Decode(&bike)
	bike.ID = params["id"]
	bikes = append(bikes, bike)
	json.NewEncoder(w).Encode(bikes)
}

// DeleteBike Delete a bike
func DeleteBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range bikes {
		if item.ID == params["id"] {
			bikes = append(bikes[:index], bikes[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(bikes)
	}
}
