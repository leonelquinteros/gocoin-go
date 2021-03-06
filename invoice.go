package gocoin

import (
	"path"
	"time"
)

// CreateInvoice is the data expected by InvoicesService.Create() method
type CreateInvoice struct {
	PriceCurrency      string `json:"price_currency,omitempty"`
	BasePrice          string `json:"base_price"`
	BasePriceCurrency  string `json:"base_price_currency"`
	Type               string `json:"type,omitempty"`
	CallbackURL        string `json:"callback_url,omitempty"`
	RedirectURL        string `json:"redirect_url,omitempty"`
	OrderID            string `json:"order_id,omitempty"`
	RefundAddress      string `json:"refund_address,omitempty"`
	ItemName           string `json:"item_name,omitempty"`
	ItemSKU            string `json:"item_sku,omitempty"`
	ItemDescription    string `json:"item_description,omitempty"`
	CustomerName       string `json:"customer_name,omitempty"`
	CustomerAddress1   string `json:"customer_address_1,omitempty"`
	CustomerAddress2   string `json:"customer_address_2,omitempty"`
	CustomerCity       string `json:"customer_city,omitempty"`
	CustomerRegion     string `json:"customer_region,omitempty"`
	CustomerCountry    string `json:"customer_country,omitempty"`
	CustomerPostalCode string `json:"customer_postal_code,omitempty"`
	CustomerEmail      string `json:"customer_email,omitempty"`
	CustomerPhone      string `json:"customer_phone,omitempty"`
	UserDefined1       string `json:"user_defined_1,omitempty"`
	UserDefined2       string `json:"user_defined_2,omitempty"`
	UserDefined3       string `json:"user_defined_3,omitempty"`
	UserDefined4       string `json:"user_defined_4,omitempty"`
	UserDefined5       string `json:"user_defined_5,omitempty"`
	UserDefined6       string `json:"user_defined_6,omitempty"`
	UserDefined7       string `json:"user_defined_7,omitempty"`
	UserDefined8       string `json:"user_defined_8,omitempty"`
	Data               string `json:"data,omitempty"`
	CustomerContact    string `json:"customer_contact,omitempty"`
	SendToCustomer     string `json:"send_to_customer,omitempty"`
}

// Invoice data
type Invoice struct {
	ID                             string            `json:"id"`
	Status                         string            `json:"status"`
	PaymentAddress                 string            `json:"payment_address"`
	Price                          string            `json:"price"`
	CryptoBalanceDue               string            `json:"crypto_balance_due"`
	PriceCurrency                  string            `json:"price_currency"`
	PriceCurrencyDetail            *CurrencyDetail   `json:"price_currency_detail"`
	ValidBillPaymentCurrencies     []string          `json:"valid_bill_payment_currencies"`
	ValidBillPaymentCurrencyDetail []*CurrencyDetail `json:"valid_bill_payment_currency_detail"`
	BasePrice                      string            `json:"base_price"`
	BasePriceCurrency              string            `json:"base_price_currency"`
	BasePriceCurrencyDetail        *CurrencyDetail   `json:"base_price_currency_detail"`
	ServiceFeeRate                 float64           `json:"service_fee_rate"`
	UsdSpotRate                    string            `json:"usd_spot_rate"`
	SpotRate                       string            `json:"spot_rate"`
	InverseSpotRate                string            `json:"inverse_spot_rate"`
	CryptoPayoutSplit              int               `json:"crypto_payout_split"`
	ConfirmationsRequired          int               `json:"confirmations_required"`
	CryptoURL                      string            `json:"crypto_url"`
	GatewayURL                     string            `json:"gateway_url"`
	NotificationLevel              string            `json:"notification_level"`
	RedirectURL                    string            `json:"redirect_url"`
	OrderID                        string            `json:"order_id"`
	ItemName                       string            `json:"item_name"`
	ItemSku                        string            `json:"item_sku"`
	ItemDescription                string            `json:"item_description"`
	Physical                       string            `json:"physical"`
	CustomerName                   string            `json:"customer_name"`
	CustomerAddress1               string            `json:"customer_address_1"`
	CustomerAddress2               string            `json:"customer_address_2"`
	CustomerCity                   string            `json:"customer_city"`
	CustomerRegion                 string            `json:"customer_region"`
	CustomerCountry                string            `json:"customer_country"`
	CustomerPostalCode             string            `json:"customer_postal_code"`
	CustomerEmail                  string            `json:"customer_email"`
	CustomerPhone                  string            `json:"customer_phone"`
	UserDefined1                   string            `json:"user_defined_1"`
	UserDefined2                   string            `json:"user_defined_2"`
	UserDefined3                   string            `json:"user_defined_3"`
	UserDefined4                   string            `json:"user_defined_4"`
	UserDefined5                   string            `json:"user_defined_5"`
	UserDefined6                   string            `json:"user_defined_6"`
	UserDefined7                   string            `json:"user_defined_7"`
	UserDefined8                   string            `json:"user_defined_8"`
	Data                           string            `json:"data"`
	ExpiresAt                      time.Time         `json:"expires_at"`
	CreatedAt                      time.Time         `json:"created_at"`
	UpdatedAt                      time.Time         `json:"updated_at"`
	ServerTime                     time.Time         `json:"server_time"`
	CallbackURL                    string            `json:"callback_url"`
	MerchantID                     string            `json:"merchant_id"`

	// ErrorResponse is embedded to catch error response bodies.
	ErrorResponse
}

// InvoicesService handles invoices API methods.
type InvoicesService struct {
	// Gocoin embeds main client
	*Gocoin
}

// Create new invoice
func (is *InvoicesService) Create(data CreateInvoice) (*Invoice, error) {
	i := new(Invoice)
	err := is.Do("POST", path.Join("merchants", is.merchantID, "invoices"), data, i)
	return i, err
}
