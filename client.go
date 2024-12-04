package coinmarketcap

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewHttpClient(endpoint string) (c *resty.Client) {
	c = resty.New().
		SetRetryCount(1).
		SetRetryWaitTime(time.Second).
		SetBaseURL(endpoint)

	return
}
