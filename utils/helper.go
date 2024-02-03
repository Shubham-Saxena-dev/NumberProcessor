package utils

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/http_client"
	"CARIAD/pkg/models"
	calculators "CARIAD/pkg/strategies"
	"net/http"
	"time"
)

func GetStrategy(req models.NumberRequest, errCollector *customerrors.ErrorCollector) calculators.NumberStrategy {
	httpClient := http.Client{Timeout: 2 * time.Second}
	client := http_client.NewHttpClient(&httpClient)
	switch req.NType {
	case models.Primes:
		return calculators.NewPrimeCalculator(calculators.NewNumberProcessor(client, *req.Url, errCollector))
	case models.Fibo:
		return calculators.NewFiboCalculator(calculators.NewNumberProcessor(client, *req.Url, errCollector))
	case models.Odd:
		return calculators.NewOddCalculator(calculators.NewNumberProcessor(client, *req.Url, errCollector))
	case models.Rand:
		return calculators.NewRandCalculator(calculators.NewNumberProcessor(client, *req.Url, errCollector))
	}
	return nil
}
