package main

import (
	"gobook/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
