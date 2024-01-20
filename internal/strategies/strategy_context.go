package calculators

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/http_client"
	"net/http"
	"net/url"
)

type NumberStrategy interface {
	ExecuteRequest() []int
	BuildRequest() *http.Request
}

type NumberProcessor struct {
	client       http_client.HTTPClient
	baseUrl      url.URL
	errCollector *customerrors.ErrorHandler
}

func NewNumberStrategy(client http_client.HTTPClient, baseUrl url.URL, collector *customerrors.ErrorHandler) NumberProcessor {
	return NumberProcessor{
		client:       client,
		baseUrl:      baseUrl,
		errCollector: collector,
	}
}
