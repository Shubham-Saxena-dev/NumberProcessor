package http_client

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Custom Http ckient")
}

var (
	mockServer   *httptest.Server
	httpClient   HTTPClient
	mockResponse interface{}
)

var _ = Describe("Http client", func() {

	BeforeSuite(func() {
		mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/success" {
				w.WriteHeader(http.StatusOK)
				requestByte, _ := json.Marshal(mockResponse)
				w.Write(requestByte)
			} else if r.URL.Path == "/error" {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}))

		httpClient = NewHttpClient(http.DefaultClient)
		mockResponse = map[string]interface{}{"myHeaderKey": "myHeaderValue"}
	})

	AfterSuite(func() {
		mockServer.Close()
	})
})
