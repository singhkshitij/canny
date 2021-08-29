package http

import (
	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func Setup() {
	client = resty.New()
}

func Get(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := client.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		Get(url)
	return resp, err
}
