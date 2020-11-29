package engine

type ErrorResponse struct {
	Code        string `json:"code,omitempty"`
	UserMessage string `json:"userMessage,omitempty"`
	DevMessage  string `json:"devMessage,omitempty"`
	HttpStatus  int    `json:"-"`
}

