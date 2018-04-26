package httperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Some code copied from https://github.com/labstack/echo/blob/master/echo.go

// WritableError is an interface for errors that can be written to http responses as JSON
type WritableError interface {
	error
	WriteJSON(w http.ResponseWriter) error
}

// HTTPError represents an http error that occurred while handling a request
type HTTPError struct {
	StatusCode int         `json:"-"`
	Message    interface{} `json:"message"`
}

// New creates a new HTTPError instance
// message can be an error, the error message is then used as message
// the message is optional but multiple messages are not supported
func New(statusCode int, message interface{}) *HTTPError {
	he := &HTTPError{StatusCode: statusCode, Message: http.StatusText(statusCode)}
	if message != nil {
		he.Message = message
	}
	return he
}

// Error makes HTTPError compatible with the error interface.
func (he *HTTPError) Error() string {
	return fmt.Sprintf("%d - %v", he.StatusCode, he.Message)
}

// WriteJSON allows to write the http error to a ResponseWriter
// for implemantation of the WritableError interface
func (he *HTTPError) WriteJSON(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(he.StatusCode)
	if err, ok := he.Message.(error); ok {
		he.Message = err.Error()
	}
	return json.NewEncoder(w).Encode(he)
}
