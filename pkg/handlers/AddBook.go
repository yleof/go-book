package handlers

import (
	"encoding/json"
	"gobook/pkg/mocks"
	"gobook/pkg/models"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book

	json.Unmarshal(body, &book)

	// Append to book mock
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// Send 201 Created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	log.Println("Book added:", book.Title)

}
