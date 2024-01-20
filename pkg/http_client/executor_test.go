package http_client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Custom HTTP Client", func() {

	Context("Process", func() {
		It("should successfully process a valid request", func() {
			request, _ := NewRequestBuilder().SetMethod(GET).SetPath(mockServer.URL+"/success").SetHeaders(HeaderContentKey, HeaderContentValue).SetBody(nil).Build()

			var response map[string]interface{}
			Ω(httpClient.BaseUrl()).To(Equal(mockServer.URL))
			result, err := httpClient.Process(request, &response)

			Ω(err).To(BeNil())
			resultMap, ok := result.(*map[string]interface{})
			Ω(resultMap).NotTo(BeNil())
			Ω(ok).To(BeTrue())
		})

		It("should handle an error response", func() {
			request, _ := NewRequestBuilder().SetMethod(GET).SetPath(mockServer.URL + "/error").SetBody(nil).Build()

			var response map[string]interface{}
			result, err := httpClient.Process(request, &response)

			Ω(result).To(BeNil())
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(ContainSubstring("500 Internal Server Error"))
		})

		It("invalid request", func() {
			mockServer.Close()
			request, _ := NewRequestBuilder().SetMethod(POST).SetPath(mockServer.URL).SetBody(nil).Build()

			var response map[string]interface{}
			result, err := httpClient.Process(request, &response)

			Ω(result).To(BeNil())
			Ω(err).To(HaveOccurred())
		})
	})

})
