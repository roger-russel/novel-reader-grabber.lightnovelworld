package helpers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const retryNumber = 13

// Download make a get to url
func Download(url string) (io.Reader, error) {
	var err error
	var req *http.Request
	var res *http.Response

	fribo := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}

	for i := 0; i < retryNumber; i++ {
		client := &http.Client{}

		req, err = http.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		//req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
		req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

		res, err = client.Do(req)

		if err != nil {
			return nil, err
		}

		if res.StatusCode == 200 {
			return res.Body, err
		}

		if res.StatusCode == http.StatusTooManyRequests {
			log.Warningf("%d too many requests while download %s, tentative number: %d", http.StatusTooManyRequests, url, i)
			time.Sleep(time.Second * time.Duration(fribo[i]))
			continue
		}

		return nil, err
	}

	return nil, fmt.Errorf("too many download tentative of %s", url)
}
