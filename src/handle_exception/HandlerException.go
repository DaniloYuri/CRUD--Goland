package handle_exception

import "net/http"

type HandlerException struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string
	Message string
}

func (h *HandlerException) Error() string {
	return h.Message
}

func NewHandlerException(message string, code int64, err string, causes []Causes) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *HandlerException {
	return &HandlerException{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
