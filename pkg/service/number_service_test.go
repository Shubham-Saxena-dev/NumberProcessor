package service

import (
	"CARIAD/pkg/models"
	"CARIAD/pkg/strategies"
	"CARIAD/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/url"
)

var _ = Describe("NumberService", func() {

	Describe("GetNumbersFromUrl", func() {
		Context("when all requests succeed", func() {
			It("should return sorted and unique numbers", func() {
				mockCacheService.Set("/primes", []int{2, 3, 5})
				mockCacheService.Set("/fibo", []int{1, 1, 2, 3, 5})
				mockCacheService.Set("/odd", []int{1, 3, 5, 7})
				numbers := numberService.GetNumbersFromUrl(requests, errCollector)
				Ω(numbers).To(Equal(utils.SortAndRemoveDuplicates([]int{1, 2, 3, 5, 7})))
			})
		})
	})

	Describe("GetTypeStrategy", func() {
		Context("when a valid request type is provided", func() {
			It("should return the corresponding number strategy", func() {
				req := models.NumberRequest{NType: models.Primes, Url: &url.URL{Path: "/primes"}}
				calculator := numberService.GetTypeStrategy(req, errCollector)
				Ω(calculator).To(BeAssignableToTypeOf(&strategies.PrimeCalculator{}))
			})
		})

		Context("when an invalid request type is provided", func() {
			It("should return nil and add an error to the error collector", func() {
				req := models.NumberRequest{NType: "invalid", Url: &url.URL{Path: "/invalid"}}
				calculator := numberService.GetTypeStrategy(req, errCollector)
				Ω(calculator).To(BeNil())
			})
		})
	})
})
