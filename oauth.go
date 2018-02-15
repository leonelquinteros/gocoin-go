package gocoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	// DashboardURL is the default Gocoin's Dashboard base URL used to obtain OAuth's Authorization Codes
	DashboardURL = "https://dashboard.gocoin.com/"
)

// AccessTokenRequest is used to exchange an authorization code for an access token.
type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
}

// AccessToken is the response format from /oauth/token method.
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`

	// OAuthErrorResponse is embedded to catch error response bodies.
	OAuthErrorResponse
}

// apiRequest wraps the API's HTTP calls using Gocoin's conventions for unauthenticated requests.
func apiRequest(method, endpoint string, data interface{}) (*http.Response, error) {
	req, err := getAPIRequest(method, endpoint, data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(req)
}

// GetAuthURL returns the OAuth URL endpoint for the client app to allow access and get a token.
func GetAuthURL(dashboardURL, merchantID, redirectURI, state string) (string, error) {
	u, err := url.Parse(dashboardURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "/auth")
	u.RawQuery = fmt.Sprintf("response_type=code&client_id=%s&redirect_uri=%s&scope=user_read&state=%s", merchantID, redirectURI, state)

	return u.String(), nil
}

// GetAccessToken exchanges an authorization code for an access token to be used on further API calls.
func GetAccessToken(apiURL, clientID, clientSecret, authCode, redirectURI string) (*AccessToken, error) {
	data := AccessTokenRequest{
		GrantType:    "authorization_code",
		Code:         authCode,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}

	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/oauth/token")

	resp, err := apiRequest("POST", u.String(), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	token := new(AccessToken)

	err = json.Unmarshal(body, token)
	return token, err
}
