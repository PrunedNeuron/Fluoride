package server

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse is a generic struct for returning a standard error document
type ErrorResponse struct {
	Error            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// ErrNotFound is a pre-built not-found error
var ErrNotFound = &ErrorResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

// Render is the Renderer for ErrResponse struct
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest is used to indicate an error on user input (with wrapped error)
func ErrInvalidRequest(err error) render.Renderer {
	var errorText string
	if err != nil {
		errorText = err.Error()
	}
	return &ErrorResponse{
		Error:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      errorText,
	}
}

// ErrorInternalLog will log an error and return a generic server error to the user
func (s *Server) ErrorInternalLog(err error) render.Renderer {
	s.logger.Errorw("Server Error", "error", err)
	return ErrorInternal(err)
}

// ErrorInternal returns a generic server error to the user
func ErrorInternal(err error) render.Renderer {
	return &ErrorResponse{
		Error:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Server Error.",
		ErrorText:      "Server Error.",
	}
}
