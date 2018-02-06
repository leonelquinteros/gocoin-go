package gocoin

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	defaultAPIURL = "https://api.gocoin.com/api/v1"
)

// apiRequest wraps the API's HTTP calls using Gocoin's conventions.
func apiRequest(method, url string, data interface{}) (*http.Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	client := &http.Client{}

	return client.Do(req)
}

// Gocoin is the main Gocoin API client.
type Gocoin struct {
	merchantID  string
	accessToken string
	apiURL      string
}

// New is the constructor for a Gocoin object.
func New(merchantID, accessToken string) *Gocoin {
	return &Gocoin{
		merchantID:  merchantID,
		accessToken: accessToken,
		apiURL:      defaultAPIURL,
	}
}

// SetAPIURL overrides the default Gocoin's API base URL and sets a custom URL.
func (g *Gocoin) SetAPIURL(apiURL string) {
	g.apiURL = apiURL
}

// SetAccessToken overrides the accessToken obtained from OAuth or from Gocoin dashboard to be used on further API calls.
func (g *Gocoin) SetAccessToken(token string) {
	g.accessToken = token
}
