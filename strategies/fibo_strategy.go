package strategies

import (
	"CARIAD/pkg/models"
	"net/http"
)

type FiboCalculator struct {
	NumberProcessor
}

func NewFiboCalculator(processor NumberProcessor) NumberStrategy {
	return &FiboCalculator{
		processor,
	}
}

func (p *FiboCalculator) ExecuteRequest() models.NumbersResponse {
	return p.NumberProcessor.ExecuteRequest()
}

func (p *FiboCalculator) BuildRequest() *http.Request {
	return p.NumberProcessor.BuildRequest()
}
