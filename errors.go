package gocoin

// ErrorResponse has the error response format from Gocoin's API.
type ErrorResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

// OAuthErrorResponse has the error response format from Gocoin's OAuth methods.
type OAuthErrorResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
