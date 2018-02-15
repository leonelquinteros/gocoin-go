package gocoin

import (
	"os"
	"testing"
)

func TestCreateInvoice(t *testing.T) {
	// Create client
	cli := New(os.Getenv("GOCOIN_MERCHANT_ID"), os.Getenv("GOCOIN_ACCESS_TOKEN"))

	// Set test data
	data := CreateInvoice{
		PriceCurrency:     "LTC",
		BasePrice:         "10.32",
		BasePriceCurrency: "USD",
		CallbackURL:       "https://testgocoin.com/callback",
	}

	invoice, err := cli.Invoices().Create(data)
	if err != nil {
		t.Fatal(err)
	}
	if invoice.Errors != nil {
		t.Errorf("%s %s %v", invoice.Status, invoice.Message, invoice.Errors)
	}
}

func TestCreateBill(t *testing.T) {
	// Create client
	cli := New(os.Getenv("GOCOIN_MERCHANT_ID"), os.Getenv("GOCOIN_ACCESS_TOKEN"))

	// Set test data
	data := CreateInvoice{
		Type:              "bill",
		BasePrice:         "10.32",
		BasePriceCurrency: "USD",
	}

	bill, err := cli.Invoices().Create(data)
	if err != nil {
		t.Fatal(err)
	}
	if bill.Errors != nil {
		t.Errorf("%s %s %v", bill.Status, bill.Message, bill.Errors)
	}
}
