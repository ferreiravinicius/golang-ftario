package controller

import "net/http"

type simpleErrorManager struct {}

func NewSimpleErrorManager() *simpleErrorManager {
	return &simpleErrorManager{}
}

func (manager simpleErrorManager) CreateResponse(error error) ErrorResponse {
	switch error.(type) {
	default:
		return createUnexpectedResponse(error)
	}
}

func createUnexpectedResponse(error error) ErrorResponse {
	code := "UnxErr"
	userMessage := "We encountered an unexpected error"
	return ErrorResponse{
		Code:        code,
		UserMessage: userMessage,
		DevMessage:  error.Error(),
		HttpStatus:  http.StatusInternalServerError,
	}
}
