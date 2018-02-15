package gocoin

import (
	"testing"
)

func TestGetAPIRequest(t *testing.T) {
	_, err := getAPIRequest("GET", "https://api.gocoin.com/api/v1/user", nil)
	if err != nil {
		t.Fatal(err)
	}
}
