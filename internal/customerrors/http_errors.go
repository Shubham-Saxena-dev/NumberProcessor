package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrHTTPClient   = errors.New("HTTP client error")
	ErrFailedToCast = errors.New("failed to cast response")
)

type ClientType string

const (
	Ren ClientType = "REN"
	Cex ClientType = "CEX"
)

type HttpError struct {
	context    string
	clientType ClientType
	payload    interface{}
	ErrType    error
	cause      error
}

func newHttpError(context string, clientType ClientType, payload interface{}, errType error, cause error) *HttpError {
	return &HttpError{
		context:    context,
		clientType: clientType,
		payload:    payload,
		ErrType:    errType,
		cause:      cause,
	}
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("%s: Client %s call failed for: %v due to %v: %v", h.context, h.clientType, h.payload, h.ErrType, h.cause)
}

func (h *HttpError) Unwrap() error {
	return h.cause
}

func ErrorHTTPClient(context string, clientType ClientType, payload interface{}, err error) *HttpError {
	return newHttpError(context, clientType, payload, ErrHTTPClient, err)
}

func ErrorFailedToCast(context string, clientType ClientType, payload interface{}, err error) *HttpError {
	return newHttpError(context, clientType, payload, ErrFailedToCast, err)
}
