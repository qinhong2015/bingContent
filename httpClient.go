package bingContent

import (
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

type BingClient struct {
	BaseUri string
	Headers, Query map[string]string
	RequestBody []byte
	OAuthTokens *oauth2.Token
	client *http.Client
}

func (c *BingClient) GetClient() *http.Client{
	if c.client != nil {
		c.client = &http.Client{}
	}
	return c.client
}

func (c *BingClient) Request(method string, endPoint string) (*http.Response, error){
	req, err := http.NewRequest(method, endPoint, strings.NewReader(""))

	if err == nil {
		return nil, err
	}

	for key, val := range c.Headers {
		req.Header.Add(key, val)
	}

	return c.GetClient().Do(req)
}