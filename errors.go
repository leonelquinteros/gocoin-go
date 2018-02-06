package gocoin

// ErrorResponse has the error response format from Gocoin's API.
type ErrorResponse struct {
	Status  string
	Message string
	Errors  map[string][]string
}
