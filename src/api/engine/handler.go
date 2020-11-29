package engine

import (
	"encoding/json"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

type AppController struct {
	errorHandler AppErrorHandler
	AppHandler
}

func NewAppController(errorManager AppErrorHandler, appHandler AppHandler) *AppController {
	return &AppController{errorHandler: errorManager, AppHandler: appHandler}
}

func (ctr AppController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	configureHeaders(w)
	if error := ctr.AppHandler(w, r); error != nil {
		response := ctr.errorHandler.CreateResponse(error)
		w.WriteHeader(response.HttpStatus)
		json.NewEncoder(w).Encode(&response) //TODO: change to encoder
	}
}

func configureHeaders(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}
