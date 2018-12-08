package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/authenticate", Authenticate).Methods("POST")
	router.HandleFunc("/bikes", AuthenticateMiddleware(GetBikes)).Methods("GET")
	router.HandleFunc("/bikes/{id}", AuthenticateMiddleware(GetBike)).Methods("GET")
	router.HandleFunc("/bikes/{id}", AuthenticateMiddleware(CreateBike)).Methods("POST")
	router.HandleFunc("/bikes/{id}", AuthenticateMiddleware(DeleteBike)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router)))
}
