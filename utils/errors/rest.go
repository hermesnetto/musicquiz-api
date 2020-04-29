package errors

import "net/http"

// RestErr defines an error related to a Rest operation
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

// NewBadRequestError creates a Bad Request error  
func NewBadRequestError(message string) * RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

// NewNotFoundtError creates a Not Found error
func NewNotFoundtError(message string) * RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}
