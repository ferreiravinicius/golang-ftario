package api

import (
	"encoding/json"
	"net/http"
)

const (
	CodeUnexpectedError = "UnxErr"
)

func handleUnexpectedError(w http.ResponseWriter, err error) {
	userMessage := "We encountered an unexpected error"
	output := ErrorResponse{
		Code:        CodeUnexpectedError,
		UserMessage: userMessage,
		DevMessage:  err.Error(),
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&output)
}