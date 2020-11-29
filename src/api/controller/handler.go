package controller

import (
	"encoding/json"
	"net/http"
)

type ErrorManager interface {
	CreateResponse(error error) ErrorResponse
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

type Route struct {
	errorManager ErrorManager
	appHandler
}

func NewRoute(errorManager ErrorManager, appHandler appHandler) *Route {
	return &Route{errorManager: errorManager, appHandler: appHandler}
}

func (route Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	configureHeaders(w)
	if error := route.appHandler(w, r); error != nil {
		response := route.errorManager.CreateResponse(error)
		w.WriteHeader(response.HttpStatus)
		json.NewEncoder(w).Encode(&response) //TODO: change to encoder
	}
}

func configureHeaders(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}

