package acceptance

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rosenhouse/jamf/application"
)

var _ = Describe("plain GET requests against a target", func() {
	var logWriter *bytes.Buffer

	BeforeEach(func() {
		logWriter = bytes.NewBuffer(nil)
	})

	It("should make a request against the given server", func() {
		var wasCalled bool
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello\n"))
			wasCalled = true
		}))
		app := application.App{TargetBaseURL: server.URL, LogWriter: logWriter}
		app.Run()
		Expect(wasCalled).To(BeTrue())
		server.Close()
	})

	Context("when no errors occur", func() {
		It("should exit status 0", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello\n"))
			}))
			app := application.App{TargetBaseURL: server.URL, LogWriter: logWriter}
			code := app.Run()
			server.Close()

			Expect(code).To(Equal(0))
		})
	})
	Context("when a connection error occurs", func() {
		It("should log the error and return code 1", func() {
			unreachableURL := "http://localhost:999999"
			app := application.App{TargetBaseURL: unreachableURL, LogWriter: logWriter}
			code := app.Run()
			Expect(code).To(Equal(1))
			Expect(logWriter.String()).To(ContainSubstring(unreachableURL))
		})
	})

	Context("when the server returns a 5xx status code", func() {
		var (
			app    application.App
			server *httptest.Server
		)

		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("server failure\n"))
			}))
			app = application.App{TargetBaseURL: server.URL, LogWriter: logWriter}
		})

		AfterEach(func() {
			server.Close()
		})

		It("should log the error", func() {
			code := app.Run()
			server.Close()

			Expect(logWriter.String()).To(ContainSubstring("500"))
			Expect(code).To(Equal(0))
		})

	})
})
