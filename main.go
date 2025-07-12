package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello, Books!")
	})

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
