package api

type ErrorResponse struct {
	Code    string `json:"code"`
	UserMessage string `json:"userMessage"`
	DevMessage string `json:"devMessage,omitempty"`
}
