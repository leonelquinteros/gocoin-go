package gocoin

import (
	"os"
	"testing"
)

func TestGetAuthURL(t *testing.T) {
	u, err := GetAuthURL(DashboardURL, os.Getenv("GOCOIN_MERCHANT_ID"), "https://testgocoin/auth/redirect", "nada")
	if err != nil {
		t.Fatal(err)
	}

	// Expected URL
	exp := "https://dashboard.gocoin.com/auth?response_type=code&client_id=" + os.Getenv("GOCOIN_MERCHANT_ID") + "&redirect_uri=https://testgocoin/auth/redirect&scope=user_read&state=nada"
	if u != exp {
		t.Errorf("Result URL not expected: %s", u)
	}
}

func TestGetAccessToken(t *testing.T) {
	// Set valid test variables here
	var (
		clientSecret = "testsecret"
		authCode     = "someauthcode"
		redirectURL  = "https://testgocoin.com/token/redirect"
	)

	// Test only if sensitive variables are set
	_, err := GetAccessToken(defaultAPIURL, os.Getenv("GOCOIN_MERCHANT_ID"), clientSecret, authCode, redirectURL)
	if err != nil {
		t.Error(err)
	}
}
