[![GitHub release](https://img.shields.io/github/release/leonelquinteros/gocoin-go.svg)](https://github.com/leonelquinteros/gocoin-go)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/leonelquinteros/gocoin-go?status.svg)](https://godoc.org/github.com/leonelquinteros/gocoin-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/gocoin-go)](https://goreportcard.com/report/github.com/leonelquinteros/gocoin-go)

# Gocoin Go SDK

Go client package for the Gocoin API (https://gocoin.com)

## Quick start

```go
package main

import (
    "fmt"
    "os"

    "github.com/leonelquinteros/gocoin"
)

func main() {
    // Create client
	cli := gocoin.New(os.Getenv("GOCOIN_MERCHANT_ID"), os.Getenv("GOCOIN_ACCESS_TOKEN"))

	// Create invoice data
	data := gocoin.CreateInvoice{
		PriceCurrency:     "LTC",
		BasePrice:         "10.32",
		BasePriceCurrency: "USD",
		CallbackURL:       "https://testgocoin.com/callback",
	}

	invoice, err := cli.Invoices().Create(data)
	if err != nil {
		panic(err)
    }
    
    fmt.Printf("Invoice generated: \n%v", invoice)
}
```

## Testing

Tests assume that there are 2 ENV variables set: 

- `GOCOIN_MERCHANT_ID` containing the MerchantID
- `GOCOIN_ACCESS_TOKEN` containing the Access Token for auth requests. 

Some tests may not use both or either at all, but most of them will need to have these set in order to work properly. 