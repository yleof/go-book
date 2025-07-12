package handlers

import (
	"encoding/json"
	"gobook/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock books
	for _, book := range mocks.Books {
		// if ids are equal, send book as response
		if book.Id == id {
			w.WriteHeader((http.StatusOK))
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}

}
