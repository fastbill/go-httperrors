# httperrors [![Build Status](https://travis-ci.org/fastbill/go-httperrors.svg?branch=master)](https://travis-ci.org/fastbill/go-httperrors) [![Go Report Card](https://goreportcard.com/badge/github.com/fastbill/go-httperrors)](https://goreportcard.com/report/github.com/fastbill/go-httperrors) [![GoDoc](https://godoc.org/github.com/fastbill/go-httperrors?status.svg)](https://godoc.org/github.com/fastbill/go-httperrors)

## Description
This package introduces a new `error` type that combines an HTTP status code and a message. It helps with two things:
1. Other packages/modules that make HTTP requests like [go-request](https://github.com/fastbill/go-request) can return this error type in case an HTTP error code was returned. The consumer of the package can then check for the type `httperrors.HTTPError` and react depending on the status code or pass on the error as is.
2. When working with frameworks like [Echo](https://echo.labstack.com) this error type can be returned in the HTTP handler and handled in the general error handler of the framework. For example you could log all 5xx errors but not invalid requests (400). Of course the error type provided by the framework could also be used for this. But the separate error type allows for a framework independent request client while still being able to just "pass on" HTTP errors if needed.

The `httperrors.HTTPError` type implements the `error` interface but also includes a `WriteJSON(w http.ResponseWriter)` method that allows to convert the error into an HTTP response. It will set the HTTP status code and write the message as JSON to the response body like this:
```javascript
{
  "message": "missing field: username"
}
```

## Example
### Creating an HTTPError
```go
func myHandler(c framework.Context) error {
	if c.FormValue("user") == "" {
		return httperrors.New(http.StatusBadRequest, "missing field: username")
	}
	
	// ...
}
```

### Handling an HTTPError
```go
func generalErrorHandler(err error, c framework.Context) {
	httpError, ok := err.(httperrors.HTTPError)
	// Alternativly a type switch can be used to identify the HTTPError.
	if ok {
	  httpError.WriteJSON(c.ResponseWriter)
	  return
	}

	// Handle other error types here ...
}
```
