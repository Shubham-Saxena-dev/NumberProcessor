package models

import (
	"CARIAD/internal/customerrors"
	"CARIAD/internal/strategies"
	"CARIAD/pkg/http_client"
	"net/http"
	"net/url"
	"time"
)

type NumberRequest struct {
	Url   *url.URL
	NType NumberType
}

func (c *NumberRequest) TypeStrategy(errCollector *customerrors.ErrorHandler) calculators.NumberStrategy {
	httpClient := http.Client{Timeout: time.Duration(20) * time.Millisecond}
	client := http_client.NewHttpClient(&httpClient)
	switch c.NType {
	case Primes:
		return calculators.NewPrimeCalculator(calculators.NewNumberStrategy(client, *c.Url, errCollector))
	case Fibo:
		return calculators.NewFiboCalculator(calculators.NewNumberStrategy(client, *c.Url, errCollector))
	case Odd:
		return calculators.NewOddCalculator(calculators.NewNumberStrategy(client, *c.Url, errCollector))
	case Rand:
		return calculators.NewRandCalculator(calculators.NewNumberStrategy(client, *c.Url, errCollector))
	}
	return nil
}
