package engine

import "net/http"

type AppErrorHandler interface {
	CreateResponse(error error) ErrorResponse
}

type DefaultErrorHandler struct {}

func NewDefaultErrorHandler() *DefaultErrorHandler {
	return &DefaultErrorHandler{}
}

func (manager DefaultErrorHandler) CreateResponse(error error) ErrorResponse {
	switch error.(type) {
	default:
		return ErrorResponse{
			UserMessage: error.Error(),
			DevMessage:  error.Error(),
			HttpStatus:  http.StatusInternalServerError,
		}
	}
}
