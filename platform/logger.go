package platform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

// Logger - api logger proxying the http request-response.
func Logger(handler func(http.ResponseWriter, *http.Request), name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// switch out response writer for a recorder
		// for all subsequent handlers
		c := httptest.NewRecorder()
		handler(c, r)

		// copy everything from response recorder
		// to actual response writer
		for k, v := range c.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		c.Body.WriteTo(w)

		// log the actual details
		dumpRequest(r)
		dumpResponse(c)

	}
}

func dumpRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))
}

func dumpResponse(w *httptest.ResponseRecorder) {
	fmt.Printf("Status code: %d\nBody: %s", w.Code, w.Body.String())
}
