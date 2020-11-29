package controller

import (
	"github.com/florestario/api/engine"
	"net/http"
)

type simpleErrorManager struct {}

func NewSimpleErrorManager() *simpleErrorManager {
	return &simpleErrorManager{}
}

func (manager simpleErrorManager) CreateResponse(error error) engine.ErrorResponse {
	switch error.(type) {
	default:
		return createUnexpectedResponse(error)
	}
}

func createUnexpectedResponse(error error) engine.ErrorResponse {
	code := "UnxErr"
	userMessage := "We encountered an unexpected error"
	return engine.ErrorResponse{
		Code:        code,
		UserMessage: userMessage,
		DevMessage:  error.Error(),
		HttpStatus:  http.StatusInternalServerError,
	}
}
