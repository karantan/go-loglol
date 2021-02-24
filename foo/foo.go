package foo

import (
	"net/http"

	"github.com/pkg/errors"
)

var httpget = http.Get

// HTTPGet function creates a request to the `url` and returns the status
func HTTPGet(url string) (int, error) {
	resp, err := httpget(url)
	if err != nil {
		log.Errorf("Error fetching URL %s. Error: %s", url, err.Error())
		return 0, errors.Wrap(err, "Error fetching URL")
	}
	log.Infof("Successfully got %s with status %s", url, resp.Status)
	return resp.StatusCode, nil
}
