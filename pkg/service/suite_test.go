package service

import (
	"CARIAD/internal/cache"
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/url"
	"testing"
)

func TestNumberService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NumberService Suite")
}

var (
	mockCacheService cache.CacheService
	errCollector     *customerrors.ErrorCollector
	numberService    NumberService
	requests         []models.NumberRequest
)

var _ = Describe("NumberService", func() {

	BeforeSuite(func() {
		mockCacheService = cache.NewCacheService()
		errCollector = customerrors.NewErrorCollector()
		numberService = NewNumberService(mockCacheService)
		requests = []models.NumberRequest{
			{NType: models.Primes, Url: &url.URL{Path: "/primes"}},
			{NType: models.Fibo, Url: &url.URL{Path: "/fibo"}},
			{NType: models.Odd, Url: &url.URL{Path: "/odd"}},
		}

		mockCacheService.Set("/primes", []int{2, 3, 5})
		mockCacheService.Set("/fibo", []int{1, 1, 2, 3, 5})
		mockCacheService.Set("/odd", []int{1, 3, 5, 7})
	})
})
