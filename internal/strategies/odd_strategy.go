package calculators

import (
	"CARIAD/pkg/http_client"
	"net/http"
)

type OddCalculator struct {
	NumberProcessor
}

func NewOddCalculator(processor NumberProcessor) NumberStrategy {
	return &OddCalculator{processor}
}

func (p *OddCalculator) ExecuteRequest() []int {
	req := p.BuildRequest()
	p.client.Process(req, []int{})
	return nil
}

func (p *OddCalculator) BuildRequest() *http.Request {
	req, err := http_client.NewRequestBuilder().
		SetMethod(http.MethodGet).
		SetPath(p.baseUrl.String()).
		SetHeaders(http_client.HeaderContentKey, http_client.HeaderContentValue).
		Build()
	if err != nil {
		p.errCollector.AddErr("primes", err)
		return nil
	}
	return req
}
