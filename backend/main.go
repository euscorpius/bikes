package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bikes", GetBikes).Methods("GET")
	router.HandleFunc("/bikes/{id}", GetBike).Methods("GET")
	router.HandleFunc("/bikes/{id}", CreateBike).Methods("POST")
	router.HandleFunc("/bikes/{id}", DeleteBike).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
