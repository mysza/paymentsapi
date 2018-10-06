package api

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Error      error  `json:"-"`      // low-level runtime error
	StatusCode int    `json:"-"`      // http response status code
	StatusText string `json:"status"` // user-level status message
}

// Render sets the application-specific error code in AppCode.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

var (
	// ErrBadRequest return status 400 Bad Request for malformed request body.
	ErrBadRequest = &ErrResponse{StatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest)}

	// ErrNotFound returns status 404 Not Found for invalid resource request.
	ErrNotFound = &ErrResponse{StatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}

	// ErrInternalServerError returns status 500 Internal Server Error.
	ErrInternalServerError = &ErrResponse{StatusCode: http.StatusInternalServerError, StatusText: http.StatusText(http.StatusInternalServerError)}
)
