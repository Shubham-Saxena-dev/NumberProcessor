package strategies

import (
	"CARIAD/pkg/models"
	"net/http"
)

type RandCalculator struct {
	NumberProcessor
}

func NewRandCalculator(processor NumberProcessor) NumberStrategy {
	return &RandCalculator{
		processor,
	}
}

func (p *RandCalculator) ExecuteRequest() models.NumbersResponse {
	return p.NumberProcessor.ExecuteRequest()
}

func (p *RandCalculator) BuildRequest() *http.Request {
	return p.NumberProcessor.BuildRequest()
}
