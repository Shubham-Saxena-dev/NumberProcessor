package strategies

import (
	"CARIAD/pkg/models"
	"net/http"
)

type PrimeCalculator struct {
	NumberProcessor
}

func NewPrimeCalculator(processor NumberProcessor) NumberStrategy {
	return &PrimeCalculator{
		processor,
	}
}

func (p *PrimeCalculator) ExecuteRequest() models.NumbersResponse {
	return p.NumberProcessor.ExecuteRequest()
}

func (p *PrimeCalculator) BuildRequest() *http.Request {
	return p.NumberProcessor.BuildRequest()
}
