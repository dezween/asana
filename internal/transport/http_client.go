package transport

import (
	"context"
	"io"
	"net/http"

	"golang.org/x/time/rate"
	"github.com/dezween/Calendar/internal/config"
)

var limiter = rate.NewLimiter(rate.Limit(config.RequestsPerSec), config.RequestsPerSec)

func MakeRequest(url string) ([]byte, error) {
	if err := limiter.Wait(context.Background()); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", config.AsanaAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
