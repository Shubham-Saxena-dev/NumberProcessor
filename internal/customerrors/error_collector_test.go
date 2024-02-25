package customerrors

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorCollector", func() {
	var errorCollector *ErrorCollector

	BeforeEach(func() {
		errorCollector = NewErrorCollector()
	})

	Describe("AddErr", func() {
		It("should add an error to the map", func() {
			err := errors.New("error occurred")
			errorCollector.AddErr("prime", err)
			Ω(errorCollector.Errors).To(HaveLen(1))
		})
	})

	Describe("Error", func() {
		It("should return the error message", func() {
			errorCollector.AddErr("a", errors.New("error 1"))
			errorCollector.AddErr("b", errors.New("error 2"))

			errMsg := errorCollector.Error()

			Ω(errMsg).To(ContainSubstring("\tError: error 1"))
			Ω(errMsg).To(ContainSubstring("\tError: error 2"))
		})
	})

})
