package middlewares

import (
	"bootcamp-web/platform/web"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Error represents an error that occurred while handling a request.
// It is usually returned by a handler to indicate that the request could not be processed
// and then processed by the error middleware.
type Error struct {
	Status  int
	Code    string
	Message string
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.Code,
		Message: e.Message,
	})
}

// StatusCode returns the http status code for the error.
// This code should be used when writing the HTTP response header as a result of handling
// this error in a middleware handler.
func (e *Error) StatusCode() int {
	return e.Status
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError creates a new error with the given status code and message.
func NewError(statusCode int, message string) error {
	return NewErrorf(statusCode, message)
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func NewErrorf(status int, format string, args ...interface{}) error {
	return &Error{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
}


// NewErrorMiddleware is a middleware that handles errors returned by handlers.
func NewErrorMiddleware() web.Middleware {
	return func(webHandler web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			err := webHandler(w, r)
			if err == nil {
				return nil
			}

			var webErr *Error
			if !errors.As(err, &webErr) {
				err = NewError(http.StatusInternalServerError, err.Error())
			}

			contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
			if m, ok := err.(json.Marshaler); ok {
				if jsonBody, marshalErr := m.MarshalJSON(); marshalErr == nil {
					contentType, body = "application/json; charset=utf-8", jsonBody
				}
			}

			w.Header().Set("Content-Type", contentType)
			if h, ok := err.(interface{ Headers() http.Header }); ok {
				for k, values := range h.Headers() {
					for _, v := range values {
						w.Header().Add(k, v)
					}
				}
			}

			code := http.StatusInternalServerError
			if sc, ok := err.(interface{ StatusCode() int }); ok {
				code = sc.StatusCode()
			}

			w.WriteHeader(code)
			_, _ = w.Write(body)

			// If any error is returned out of the middleware stack, the application will
			// panic. So, we return nil here to indicate that the error has been
			// handled.
			return nil
		}
	}
}
