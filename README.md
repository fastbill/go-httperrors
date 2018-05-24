# httperrors [![Build Status](https://travis-ci.org/fastbill/httperrors.svg?branch=master)](https://travis-ci.org/fastbill/httperrors) [![Go Report Card](https://goreportcard.com/badge/github.com/fastbill/httperrors)](https://goreportcard.com/report/github.com/fastbill/httperrors) [![GoDoc](https://godoc.org/github.com/fastbill/httperrors?status.svg)](https://godoc.org/github.com/fastbill/httperrors)

This package introduces a type that combines status code and response. 
It's designed for frameworks like [Echo](https://echo.labstack.com) where every handler returns an `error` and an error handler will check what to do after an handler returned one.

This allows elegant code like this:

```go
func myHandler(c framework.Context) error {
	if c.FormValue("user") == "" {
		return httperrors.New(http.StatusBadRequest, "missing field: username")
	}
	
	// ...
}
```

The error handler will then have to typecheck against `httperrors.HTTPError` and if it's one, 
call `HTTPError.WriteJSON(w)`. Otherwise just log it and return `Internal Server Error`.
