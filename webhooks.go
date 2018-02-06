package gocoin

// Webhook is the type that contain generic webhooks bodies
type Webhook struct {
	ID      string      `json:"id"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}
