package service

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/models"
	"sync"
)

type NumberService interface {
	GetNumbersFromUrl([]models.NumberRequest) []models.NumbersResponse
}

type numbersService struct {
	cache        map[string][]int
	errCollector *customerrors.ErrorHandler
}

func NewNumberService(errCollector *customerrors.ErrorHandler) NumberService {
	return &numbersService{
		cache:        make(map[string][]int),
		errCollector: errCollector,
	}
}

func (n *numbersService) GetNumbersFromUrl(reqs []models.NumberRequest) []models.NumbersResponse {
	var responses []models.NumbersResponse
	wg := sync.WaitGroup{}
	numberCh := make(chan []int, len(reqs))
	for _, req := range reqs {
		calculator := req.TypeStrategy(n.errCollector)
		if calculator == nil {
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			numbers := calculator.ExecuteRequest()
			numberCh <- numbers
		}(&wg)
	}
	go func() {
		wg.Wait()
		close(numberCh)
	}()
	for numbers := range numberCh {
		responses = append(responses, models.NumbersResponse{Numbers: numbers})
	}

	return responses
}
