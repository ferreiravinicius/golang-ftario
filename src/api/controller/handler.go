package controller

import (
	"encoding/json"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

const contentTypeKey, contentTypeValue = "Content-Type", "application/json"

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)
	if error := handler(w, r); error != nil {
		response := ManageErrorResponse(error)
		w.WriteHeader(response.HttpStatus)
		json.NewEncoder(w).Encode(&response) //TODO: change to encoder
	}
}
