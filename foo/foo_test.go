package foo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	description string
	input       int
	expected    int
}{
	{
		description: "200 response",
		input:       200,
		expected:    200,
	},
	{
		description: "404 response",
		input:       404,
		expected:    404,
	},
	{
		description: "503 response",
		input:       503,
		expected:    503,
	},
}

func TestHTTPGet(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			server := makeServer(tc.input)
			got, _ := HTTPGet(server.URL)
			assert.Equal(got, tc.expected, "should be equal")
		})
	}
}

func BenchmarkHTTPGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			server := makeServer(tc.input)
			HTTPGet(server.URL)
		}
	}
}

func makeServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
	}))
}
