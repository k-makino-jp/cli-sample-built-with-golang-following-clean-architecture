// http is an package of HTTP request
package http

import (
	"github.com/go-resty/resty/v2"
)

type client struct {
	restyClient *resty.Client
}

func NewClient() *client {
	return &client{
		restyClient: resty.New(),
	}
}

// Get executes HTTP request with GET method.
func (c client) Get(queries map[string]string, headers map[string]string, url string) (statusCode int, respBody []byte, err error) {
	resp, err := c.restyClient.R().
		SetQueryParams(queries).
		SetHeaders(headers).
		Get(url)
	return resp.StatusCode(), resp.Body(), err
}

// Post executes HTTP request with POST method.
func (c client) Post(queries map[string]string, headers map[string]string, body string, url string) (statusCode int, respBody []byte, err error) {
	resp, err := c.restyClient.R().
		SetQueryParams(queries).
		SetHeaders(headers).
		SetBody(body).
		Post(url)
	return resp.StatusCode(), resp.Body(), err
}
