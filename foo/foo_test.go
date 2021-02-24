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
		description: "200 OK",
		input:       200,
		expected:    200,
	},
	{
		description: "404 Not Found",
		input:       404,
		expected:    404,
	},
	{
		description: "503 Internal Server Error",
		input:       503,
		expected:    503,
	},
}

// // Test HTTPGet function by creating a http test server
func TestHTTPGetMockHTTPServer(t *testing.T) {
	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			server := makeServer(tc.input)
			got, _ := HTTPGet(server.URL)
			assert.Equal(got, tc.expected, "should be equal")
		})
	}
}

// Test HTTPGet function with overloading the `httpget` variable function
func TestHTTPGetVariableOverload(t *testing.T) {
	assert := assert.New(t)
	for _, tc := range testCases {
		httpget = func(url string) (resp *http.Response, err error) {
			resp = &http.Response{
				Status:     tc.description,
				StatusCode: tc.expected,
			}
			return resp, err
		}
		t.Run(tc.description, func(t *testing.T) {
			got, _ := HTTPGet("foo")
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
