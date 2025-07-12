package handlers

import (
	"encoding/json"
	"gobook/pkg/mocks"
	"gobook/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	// iterate over all the mock books
	for i, book := range mocks.Books {
		// if ids are equal, update the book
		if book.Id == id {
			// Update the book in the mock
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc
			mocks.Books[i] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			log.Println("Book updated:", updatedBook.Title)
			break
		}
	}

}
