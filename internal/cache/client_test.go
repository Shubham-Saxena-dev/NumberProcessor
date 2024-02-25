package cache

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNumberService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cache service Suite")
}

var _ = Describe("CacheService", func() {
	var (
		cacheService CacheService
		path         string
		numbers      []int
	)

	BeforeSuite(func() {
		cacheService = NewCacheService()
		path = "/prime"
		numbers = []int{1, 2, 3}
		cacheService.Set(path, numbers)
	})

	Describe("Get", func() {
		Context("when the path exists in the cache", func() {
			It("should return cached numbers and true", func() {
				cachedNumbers, exists := cacheService.Get(path)
				Ω(exists).To(BeTrue())
				Ω(cachedNumbers).To(Equal(numbers))
			})
		})

		Context("when the path does not exist in the cache", func() {
			It("should return nil and false", func() {
				path := "/abc"
				cachedNumbers, exists := cacheService.Get(path)
				Ω(exists).To(BeFalse())
				Ω(cachedNumbers).To(BeNil())
			})
		})
	})

	Describe("Set", func() {
		It("should add or update the numbers for the given path in the cache", func() {
			path := "/odd"
			numbers := []int{2, 3, 7}
			cacheService.Set(path, numbers)
			cachedNumbers, exists := cacheService.Get(path)
			Ω(exists).To(BeTrue())
			Ω(cachedNumbers).To(Equal(numbers))
		})
	})

	Describe("Delete", func() {
		Context("when the path exists in the cache", func() {
			It("should delete the path from the cache", func() {
				path := "/prime"
				numbers := []int{1, 2, 3}
				cacheService.Set(path, numbers)
				cacheService.Delete(path)

				cachedNumbers, exists := cacheService.Get(path)
				Ω(exists).To(BeFalse())
				Ω(cachedNumbers).To(BeNil())
			})
		})

		Context("when the path does not exist in the cache", func() {
			It("should do nothing", func() {
				path := "/fibo"
				cacheService.Delete(path)
				cachedNumbers, exists := cacheService.Get(path)
				Ω(exists).To(BeFalse())
				Ω(cachedNumbers).To(BeNil())
			})
		})
	})
})
