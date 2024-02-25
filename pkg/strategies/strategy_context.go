package strategies

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/http_client"
	"CARIAD/pkg/models"
	"net/http"
	"net/url"
)

type NumberStrategy interface {
	ExecuteRequest() models.NumbersResponse
	BuildRequest() *http.Request
}

type NumberProcessor struct {
	client       http_client.HTTPClient
	baseUrl      url.URL
	errCollector *customerrors.ErrorCollector
}

func NewNumberProcessor(client http_client.HTTPClient, baseUrl url.URL, collector *customerrors.ErrorCollector) NumberProcessor {
	return NumberProcessor{
		client:       client,
		baseUrl:      baseUrl,
		errCollector: collector,
	}
}

func (p *NumberProcessor) ExecuteRequest() models.NumbersResponse {
	req := p.BuildRequest()
	var response models.NumbersResponse
	process, err := p.client.Process(req, &response)
	if err != nil || process == nil {
		p.errCollector.AddErr(p.baseUrl.String(), err)
		return models.EmptyResponse
	}

	return *process.(*models.NumbersResponse)
}

func (p *NumberProcessor) BuildRequest() *http.Request {
	req, err := http_client.NewRequestBuilder().
		SetMethod(http.MethodGet).
		SetPath(p.baseUrl.String()).
		Build()
	if err != nil {
		p.errCollector.AddErr(p.baseUrl.String(), err)
		return nil
	}
	return req
}
