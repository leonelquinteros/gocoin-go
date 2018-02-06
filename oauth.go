package gocoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

// GetAuthURL returns the OAuth URL endpoint for the client app to allow access and get a token.
func GetAuthURL(dashboardURL, merchantID, redirectURI, state string) string {
	url := path.Join(dashboardURL, "/auth?response_type=code&client_id=%s&redirect_uri=%s&scope=user_read&state=%s")

	return fmt.Sprintf(url, merchantID, redirectURI, state)
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

	resp, err := apiRequest("POST", path.Join(apiURL, "/oauth/token"), data)
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
