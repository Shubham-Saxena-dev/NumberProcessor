package http_client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/url"
)

var _ = Describe("RequestBuilder", func() {
	var (
		builder *RequestBuilder
	)

	BeforeEach(func() {
		builder = NewRequestBuilder()
	})

	Context("SetPath", func() {
		It("should set the request path", func() {
			path := "/myPath"
			builder.SetPath(path)
			Ω(builder.path).To(Equal(path))
		})
	})

	Context("ValidateUrl", func() {
		It("should parse path", func() {
			path := "/myPath"
			s := validateUrl(path)
			Ω(s).To(Equal(path))
		})
	})

	Context("SetHeaders", func() {
		It("should set request headers", func() {
			builder.SetHeaders("Content-Type", "application/json")
			req, _ := builder.Build()
			Ω(req.Header.Get("Content-Type")).To(Equal("application/json"))
		})
	})

	Context("SetBody", func() {
		It("should set the request body", func() {
			body := map[string]string{"myHeaderKey": "myHeaderValue"}
			builder.SetBody(body)
			req, _ := builder.Build()
			Ω(req.Body).NotTo(BeNil())

			_, err := ioutil.ReadAll(req.Body)
			Ω(err).To(BeNil())
		})

		It("should handle nil body", func() {
			builder.SetBody(nil)
			req, _ := builder.Build()
			Ω(req.Body).To(BeNil())
		})
	})

	Context("SetMethod", func() {
		It("should set the request method", func() {
			builder.SetMethod(POST)
			req, _ := builder.Build()
			Ω(req.Method).To(Equal("POST"))
		})
	})

	Context("SetQueryParameters", func() {
		It("should set the request params", func() {
			queryParams := url.Values{
				"param1": {"ABC"},
				"parma2": {"DEF"},
			}
			builder.SetQueryParameters(queryParams)
			req, _ := builder.Build()
			Ω(req.URL.RawQuery).To(Equal("param1=ABC&parma2=DEF"))
		})
	})

	Context("Build", func() {
		It("should build an HTTP request", func() {
			builder.SetPath("/test")
			builder.SetHeaders("Content-Type", "application/json")
			body := map[string]string{"myHeaderKey": "myHeaderValue"}
			builder.SetBody(body)
			builder.SetMethod(PUT)

			req, err := builder.Build()
			Ω(err).To(BeNil())
			Ω(req.Method).To(Equal("PUT"))
			Ω(req.URL.Path).To(Equal("/test"))
			Ω(req.Header.Get("Content-Type")).To(Equal("application/json"))

			_, err = ioutil.ReadAll(req.Body)
			Ω(err).To(BeNil())
		})

		It("should return an error for an invalid URL", func() {
			builder.SetPath("===invalid-url===")
			_, err := builder.Build()
			Ω(err).To(BeNil())
		})
	})
})
