package gocoin

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	defaultAPIURL = "https://api.gocoin.com/api/v1"
)

// getAPIRequest creates a http.Request object using Gocoin's conventions
func getAPIRequest(method, url string, data interface{}) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	if data != nil {
		// Create request with body data
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	} else {
		// Create empty body request
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	return req, nil
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
// This method isn't safe to be used concurrently.
func (g *Gocoin) SetAPIURL(apiURL string) {
	g.apiURL = apiURL
}

// SetAccessToken overrides the accessToken obtained from OAuth or from Gocoin dashboard to be used on further API calls.
// This method isn't safe to be used concurrently.
func (g *Gocoin) SetAccessToken(token string) {
	g.accessToken = token
}

// APIRequest makes low level Gocoin API requests using object's configuration
func (g *Gocoin) APIRequest(method, url string, data interface{}) (*http.Response, error) {
	req, err := getAPIRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+g.accessToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// Do performs an API request and unmarshal the response body into the provided object pointer
func (g *Gocoin) Do(method, url string, data, response interface{}) error {
	resp, err := g.APIRequest(method, url, data)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, response)
}

// User returns a UserService object associated to this client
func (g *Gocoin) User() *UserService {
	return &UserService{
		Gocoin: g,
	}
}

// Invoices returns a InvoicesService object associated to this client
func (g *Gocoin) Invoices() *InvoicesService {
	return &InvoicesService{
		Gocoin: g,
	}
}
