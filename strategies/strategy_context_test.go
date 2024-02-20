package strategies

import (
	"CARIAD/pkg/http_client"
	"CARIAD/pkg/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/url"
)

var _ = Describe("NumberProcessor", func() {

	It("should execute request and return numbers response", func() {
		baseUrl, _ := url.Parse(mockServer.URL)
		mockHTTPClient = http.DefaultClient
		numberProcessor := NewNumberProcessor(
			http_client.NewHttpClient(mockHTTPClient),
			*baseUrl,
			errCollector,
		)

		response := numberProcessor.ExecuteRequest()
		Ω(response).To(Equal(models.NumbersResponse{Numbers: []int{1, 2, 3}}))
	})

	It("should handle error when executing request", func() {
		mockHTTPClient = http.DefaultClient
		numberProcessor := NewNumberProcessor(
			http_client.NewHttpClient(http.DefaultClient),
			url.URL{Path: "/fail"},
			errCollector,
		)
		response := numberProcessor.ExecuteRequest()
		Ω(errCollector.Errors).To(HaveLen(1))
		Ω(response).To(Equal(models.EmptyResponse))
	})
})
