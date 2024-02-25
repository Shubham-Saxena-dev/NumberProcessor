package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrHTTPClient     = errors.New("HTTP client error")
	ErrFailedToDecode = errors.New("failed to decode into given response")
)

type HttpError struct {
	context string
	payload interface{}
	ErrType error
	cause   error
}

func newHttpError(context string, errType error, cause error) *HttpError {
	return &HttpError{
		context: context,
		ErrType: errType,
		cause:   cause,
	}
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("%s: client call failed due to %v: %v", h.context, h.ErrType, h.cause)
}

func (h *HttpError) Unwrap() error {
	return h.cause
}

func ErrorHTTPClient(context string, err error) *HttpError {
	return newHttpError(context, ErrHTTPClient, err)
}

func ErrorFailedToDecode(context string, err error) *HttpError {
	return newHttpError(context, ErrFailedToDecode, err)
}
