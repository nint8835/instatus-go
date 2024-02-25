package instatus_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (c *Client) get(url *url.URL, target any) error {
	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var e Error
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return fmt.Errorf("error decoding error response: %w", err)
		}

		return e
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}

func New(key string) *Client {
	return &Client{
		httpClient: &http.Client{
			Transport: &authenticatedRoundtripper{key: key},
		},
	}
}
