package errors

import "net/http"

type RestError struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func NewBadRequestError(m string) *RestError {
	return &RestError{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   http.StatusText(http.StatusBadRequest),
	}
}

func NewInternalRequestError(m string) *RestError {
	return &RestError{
		Message: m,
		Status:  http.StatusInternalServerError,
		Error:   http.StatusText(http.StatusInternalServerError),
	}
}

func OK(m interface{}) *RestError {
	return &RestError{
		Data:   m,
		Status: http.StatusOK,
		Error:  http.StatusText(http.StatusOK),
	}
}

func NewNotFoundError(m interface{}) *RestError {
	return &RestError{
		Data:   m,
		Status: http.StatusNotFound,
		Error:  http.StatusText(http.StatusNotFound),
	}
}
