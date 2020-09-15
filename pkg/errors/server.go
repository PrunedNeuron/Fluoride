package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse represents the error response object
type ErrorResponse struct {
	Error          error  `json:"-"`               // low level runtime err
	HTTPStatusCode int    `json:"-"`               // http status code
	Message        string `json:"message"`         // high level status message
	Status         string `json:"status"`          // status message (success | failure)
	ErrorCode      int    `json:"code,omitempty"`  // application level error code
	ErrorText      string `json:"error,omitempty"` // application level error message
}

// ErrNotFound is a predefined 404 error
var ErrNotFound = &ErrorResponse{
	HTTPStatusCode: 404,
	Status:         "failure",
	Message:        "resource not found",
}

// Render is the error response renderer
func (err *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, err.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest indicates that the request was invalid
func ErrInvalidRequest(err error) render.Renderer {
	var errorText string
	if err != nil {
		errorText = err.Error()
	}

	return &ErrorResponse{
		Error:          err,
		HTTPStatusCode: http.StatusBadRequest,
		Status:         "failure",
		Message:        "invalid request",
		ErrorText:      errorText,
	}
}

// ErrInternalServer returns a generic server error
func ErrInternalServer(err error) render.Renderer {

	var errorText string

	if err != nil {
		errorText = err.Error()
	}
	return &ErrorResponse{
		Error:          err,
		HTTPStatusCode: http.StatusInternalServerError,
		Status:         "failure",
		Message:        "server error",
		ErrorText:      errorText,
	}
}
