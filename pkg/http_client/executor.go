package http_client

import (
	"CARIAD/internal/customerrors"
	"fmt"
	"github.com/avast/retry-go"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	NumberOfRetries = 3
	DelayPeriod     = 10 * time.Millisecond
)

type HTTPClient interface {
	Process(*http.Request, interface{}) (interface{}, error)
}

type client struct {
	httpClient *http.Client
}

func NewHttpClient(httpClient *http.Client) HTTPClient {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) Process(request *http.Request, response interface{}) (interface{}, error) {
	resp, err := executeRequest(request, c.httpClient)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, customerrors.ErrorHTTPClient(request.URL.Path, err)
	}

	if err := jsoniter.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, customerrors.ErrorFailedToDecode(request.URL.Path, err)
	}
	return response, nil
}

func executeRequest(request *http.Request, httpClient *http.Client) (*http.Response, error) {
	var resp *http.Response
	err := retry.Do(
		func() error {
			var err error
			resp, err = httpClient.Do(request)
			return err
		},
		retry.Attempts(NumberOfRetries),
		retry.Delay(DelayPeriod),
		retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("Retrying request for %d times because of error: %v", n+1, err))
		}),
	)
	return resp, err
}
