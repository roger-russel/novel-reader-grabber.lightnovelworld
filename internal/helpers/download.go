package helpers

import (
	"fmt"
	"io"
	"net/http"
)

// Download make a get to url
func Download(url string) (io.Reader, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, err

}
