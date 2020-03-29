package client

import (
	"fmt"
	"net/http"
)

type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Response *http.Response
	Body []byte
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unauthorized error: %v\nResponse: %s", e.Response.Status, e.Body)
}

type ApiError struct {
	Response *http.Response
	Body []byte
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Api error: %v\nResponse: %s", e.Response.Status, e.Body)
}