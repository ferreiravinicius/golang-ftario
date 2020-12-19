package engine

import (
	"encoding/json"
	"net/http"
)

type AppHttpHandler func(w http.ResponseWriter, r *http.Request) error

type AppContext struct {
	errorHandler AppErrorHandler
	AppHttpHandler
}

func NewAppContext(errorHandler AppErrorHandler, httpHandler AppHttpHandler) *AppContext {
	return &AppContext{errorHandler: errorHandler, AppHttpHandler: httpHandler}
}

func (ctx AppContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	configureHeaders(&w)
	if error := ctx.AppHttpHandler(w, r); error != nil {
		response := ctx.errorHandler.CreateResponse(error)
		w.WriteHeader(response.HttpStatus)
		json.NewEncoder(w).Encode(&response) //TODO: change to encoder
	}
}

func configureHeaders(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Add("Access-Control-Allow-Origin", "*")
	(*responseWriter).Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	(*responseWriter).Header().Add("Access-Control-Allow-Headers", "Content-Type")
	(*responseWriter).Header().Set("Content-Type", "application/json")
}
