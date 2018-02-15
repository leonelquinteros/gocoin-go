package gocoin

import (
	"os"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Create client
	cli := New(os.Getenv("GOCOIN_MERCHANT_ID"), os.Getenv("GOCOIN_ACCESS_TOKEN"))

	u, err := cli.User().Get()
	if err != nil {
		t.Fatal(err)
	}
	if u.Errors != nil {
		t.Errorf("%s %s %v", u.Status, u.Message, u.Errors)
	}
}
