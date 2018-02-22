package gocoin

import (
	"time"
)

// User object represent responses from /api/v1/user method
type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ImageURL   string    `json:"image_url"`
	MerchantID string    `json:"merchant_id"`
	Confirmed  bool      `json:"confirmed,omitempty"`
	Roles      []string  `json:"roles,omitempty"`

	// ErrorResponse is embedded to catch error response bodies.
	ErrorResponse
}

// UserService handles User API methods
type UserService struct {
	// Gocoin embeds main client
	*Gocoin
}

// Get calls GET /api/v1/user and returns a User object
func (us *UserService) Get() (*User, error) {
	u := new(User)
	err := us.Do("GET", "/user", nil, u)
	return u, err
}
