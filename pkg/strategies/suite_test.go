package strategies

import (
	"CARIAD/internal/customerrors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculators Suite")
}

var (
	mockHTTPClient *http.Client
	mockServer     *httptest.Server
	errCollector   *customerrors.ErrorCollector
)

var _ = Describe("NumberProcessor", func() {
	BeforeSuite(func() {
		mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/error" {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"Numbers": [1, 2, 3]}`))
			}
		}))
		errCollector = customerrors.NewErrorCollector()
	})

	AfterSuite(func() {
		mockServer.Close()
	})
})
