package customerrors

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorCollector", func() {
	Describe("HttpErrors", func() {
		It("should create a httpError and format the error message correctly", func() {
			err := errors.New("error occurred")
			httpErr := ErrorHTTPClient("a", err)
			Ω(httpErr.Error()).To(Equal("a: client call failed due to HTTP client error: error occurred"))
			httpErr = ErrorFailedToDecode("b", err)
			Ω(httpErr.Error()).To(Equal("b: client call failed due to failed to decode into given response: error occurred"))
			Ω(httpErr.Unwrap().Error()).To(Equal(err.Error()))
		})
	})
})
