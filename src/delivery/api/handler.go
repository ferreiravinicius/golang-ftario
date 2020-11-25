package api

import (
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

const (
	contentTypeKey, contentTypeValue = "Content-Type", "application/json"
)

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)
	if err := handler(w, r); err != nil {
		switch e := err.(type) {
		default:
			handleUnexpectedError(w, e)
		}
	}
}
