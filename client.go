package instatus_go

import (
	"net/http"
)

type authenticatedRoundtripper struct {
	key string
}

func (a *authenticatedRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.key)
	return http.DefaultTransport.RoundTrip(req)
}

type Client struct {
	httpClient *http.Client
}

func New(key string) *Client {
	return &Client{
		httpClient: &http.Client{
			Transport: &authenticatedRoundtripper{key: key},
		},
	}
}
