package gocoin

// Webhook is the type that contain generic webhooks bodies
type Webhook struct {
	ID      string      `json:"id"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

// InvoiceWebhook is the type that contain the expected invoice webhook payload
type InvoiceWebhook struct {
	ID      string  `json:"id"`
	Event   string  `json:"event"`
	Payload Invoice `json:"payload"`
}
