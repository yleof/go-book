package main

import (
	"gobook/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET") // Assuming you want to handle GET by ID as well
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
