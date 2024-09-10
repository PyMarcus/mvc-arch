package errs

import (
	"net/http"
)

// Errs takes message: the error description, err: stacktrace
type Errs struct {
	Message string    `json:"message"`
	Err     string    `json:"error"`
	Code    int     `json:"code"`
	Causes  []*Causes `json:"causes"`
}

func (e Errs) Error() string {
	return e.Message
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewErrs(message, err string, code int, causes []*Causes) *Errs {
	return &Errs{
		message,
		err,
		code,
		causes,
	}
}

func NewBadRequestError(message string) *Errs {
	return &Errs{
		message,
		"bad_request",
		http.StatusBadRequest,
		nil,
	}
}

func NewBadRequestValidationError(message string, causes []*Causes) *Errs {
	return &Errs{
		message,
		"bad_request",
		http.StatusBadRequest,
		causes,
	}
}

func NewInternalServerError(message string) *Errs {
	return &Errs{
		message,
		"internal_server_error",
		http.StatusInternalServerError,
		nil,
	}
}

func NewNotFoundError(message string) *Errs {
	return &Errs{
		message,
		"not_found",
		http.StatusNotFound,
		nil,
	}
}

func NewForbiddenError(message string) *Errs {
	return &Errs{
		message,
		"forbidden",
		http.StatusForbidden,
		nil,
	}
}
