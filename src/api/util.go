package api

import (
	"encoding/json"
	"net/http"
)

type ErrorOutput struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func AnswerError(w http.ResponseWriter, statusCode int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	output := ErrorOutput{
		Code:    code,
		Message: message,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&output)
}

func AnswerSuccess(w http.ResponseWriter, statusCode int, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if len(data) == 1 {
		json.NewEncoder(w).Encode(&data)
	}
}
