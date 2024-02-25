package strategies

import (
	"CARIAD/pkg/models"
	"net/http"
)

type OddCalculator struct {
	NumberProcessor
}

func NewOddCalculator(processor NumberProcessor) NumberStrategy {
	return &OddCalculator{processor}
}

func (p *OddCalculator) ExecuteRequest() models.NumbersResponse {
	return p.NumberProcessor.ExecuteRequest()
}

func (p *OddCalculator) BuildRequest() *http.Request {
	return p.NumberProcessor.BuildRequest()
}
