package service

import (
	"CARIAD/internal/cache"
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/models"
	calculators "CARIAD/pkg/strategies"
	"CARIAD/utils"
	"sync"
	"time"
)

type NumberService interface {
	GetNumbersFromUrl([]models.NumberRequest, *customerrors.ErrorCollector) []int
	GetTypeStrategy(models.NumberRequest, *customerrors.ErrorCollector) calculators.NumberStrategy
}

type numbersService struct {
	cacheService cache.CacheService
	errCollector *customerrors.ErrorCollector
	mu           sync.Mutex
}

func NewNumberService(cacheService cache.CacheService) NumberService {
	return &numbersService{
		cacheService: cacheService,
	}
}

func (n *numbersService) GetNumbersFromUrl(reqs []models.NumberRequest, errCollector *customerrors.ErrorCollector) []int {
	var responses []int
	var wg sync.WaitGroup
	done := make(chan struct{})

	for _, req := range reqs {
		wg.Add(1)
		go func(req models.NumberRequest) {
			defer wg.Done()

			calculator := n.GetTypeStrategy(req, errCollector)
			if calculator == nil {
				return
			}

			val, exists := n.cacheService.Get(req.Url.Path)
			if exists {
				n.mu.Lock()
				responses = append(responses, val...)
				n.mu.Unlock()
				return
			}

			numbers := calculator.ExecuteRequest()
			if numbers.Numbers != nil {
				n.mu.Lock()
				n.cacheService.Set(req.Url.Path, numbers.Numbers)
				responses = append(responses, numbers.Numbers...)
				n.mu.Unlock()
			}
		}(req)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return utils.SortAndRemoveDuplicates(responses)
	case <-time.After(500 * time.Second):
		return utils.SortAndRemoveDuplicates(responses)
	}
}

func (n *numbersService) GetTypeStrategy(req models.NumberRequest, errCollector *customerrors.ErrorCollector) calculators.NumberStrategy {
	return utils.GetStrategy(req, errCollector)
}
