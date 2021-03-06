package client

import (
	"fmt"
	"log"

	"gopkg.in/resty.v1"
)

var client = resty.New()

// NewHTTPRequest create http request
func NewHTTPRequest() *resty.Request {
	return client.R()
}

func responseLogger(c *resty.Client, resp *resty.Response) error {
	method := fmt.Sprintf("[Method] %s", resp.Request.Method)
	url := fmt.Sprintf("[URL] %s", resp.Request.URL)
	reqBody := fmt.Sprintf("[Request Body] %v", resp.Request.Body)
	status := fmt.Sprintf("[Status] %d", resp.StatusCode())
	duration := fmt.Sprintf("[Duration] %v", resp.Time())
	respBody := fmt.Sprintf("[Response Body] %s", resp.String())

	log := fmt.Sprintf("%s %s %s %s %s %s", method, url, reqBody, status, duration, respBody)
	c.Log.Println(log)
	return nil
}

// Setup setup http client
func Setup(log *log.Logger) {
	client.OnAfterResponse(responseLogger)
	client.Log = log
}
