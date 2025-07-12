package handlers

import (
	"encoding/json"
	"gobook/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock books
	for i, book := range mocks.Books {
		// if ids are equal, delete the book
		if book.Id == id {
			mocks.Books = append(mocks.Books[:i], mocks.Books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}

}
