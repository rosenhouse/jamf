package acceptance

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rosenhouse/jamf/application"
)

var _ = Describe("Making a request against a server", func() {
	var logWriter *bytes.Buffer

	BeforeEach(func() {
		logWriter = bytes.NewBuffer(nil)
	})

	It("makes a request against the given server", func() {
		var wasCalled bool
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello\n"))
			wasCalled = true
		}))

		app := application.App{
			TargetBaseURL: server.URL,
		}

		app.Run()

		Expect(wasCalled).To(BeTrue())

		server.Close()
	})

	Context("when an connection error occurs", func() {
		It("logs the error and returns non-zero", func() {
			unreachableURL := "http://localhost:999999"
			app := application.App{TargetBaseURL: unreachableURL, LogWriter: logWriter}

			code := app.Run()

			Expect(code).NotTo(Equal(0))
			Expect(logWriter.String()).To(ContainSubstring(unreachableURL))
		})
	})
})
