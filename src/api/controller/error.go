package controller

import "net/http"

// Creates ErrorResponse based on error type
func ManageErrorResponse(error error) ErrorResponse {
	switch error.(type) {
	default:
		return createUnexpectedResponse(error)
	}
}

const (
	CodeUnexpectedError = "UnxErr"
)

func createUnexpectedResponse(error error) ErrorResponse {
	userMessage := "We encountered an unexpected error"
	return ErrorResponse{
		Code:        CodeUnexpectedError,
		UserMessage: userMessage,
		DevMessage:  error.Error(),
		HttpStatus:  http.StatusInternalServerError,
	}
}
