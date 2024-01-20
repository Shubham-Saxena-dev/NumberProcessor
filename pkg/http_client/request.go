package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

const (
	HeaderContentKey   = "Content-Type"
	HeaderContentValue = "application/json"
)

type RequestBuilder struct {
	path        string
	headers     map[string][]string
	body        io.ReadCloser
	method      Method
	ctx         context.Context
	queryParams url.Values
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		method:  GET,
		headers: map[string][]string{},
		ctx:     context.Background(),
		body:    nil,
	}
}

func (b *RequestBuilder) SetPath(path string) *RequestBuilder {
	b.path = validateUrl(path)
	return b
}

func (b *RequestBuilder) SetHeaders(name, value string) *RequestBuilder {
	values, found := b.headers[name]
	if !found {
		values = make([]string, 0, 10)
	}
	b.headers[name] = append(values, value)
	return b
}

func (b *RequestBuilder) SetBody(body interface{}) *RequestBuilder {
	if body != nil {
		bodyJSON, err := json.Marshal(b.body)
		if err != nil {
			log.Error(fmt.Sprintf("unable to marshall request body with error : %v", err))
		}
		b.body = io.NopCloser(bytes.NewBuffer(bodyJSON))
	}
	return b
}

func (b *RequestBuilder) SetMethod(method Method) *RequestBuilder {
	b.method = method.validateMethodType()
	return b
}

func validateUrl(path string) string {
	serverURL, err := url.Parse(path)
	if err != nil {
		return ""
	}
	return serverURL.String()
}

func (b *RequestBuilder) SetQueryParameters(params url.Values) *RequestBuilder {
	b.queryParams = params
	return b
}

func (b *RequestBuilder) Build() (*http.Request, error) {
	if len(b.queryParams) > 0 {
		b.path += "?" + b.queryParams.Encode()
	}
	req, err := http.NewRequestWithContext(b.ctx, string(b.method), b.path, b.body)
	if err != nil {
		return nil, err
	}

	for key, values := range b.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}
