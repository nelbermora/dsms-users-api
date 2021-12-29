package model

type StandardError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewStError(code int, message string) StandardError {
	return StandardError{
		Code:    code,
		Message: message,
	}
}
