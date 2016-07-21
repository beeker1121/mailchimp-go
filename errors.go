package mailchimp

import (
	"errors"
	"fmt"
)

var (
	// ErrAPIKeyNotSet is returned when a call to the API is
	// attempted before the user set an API key.
	ErrAPIKeyNotSet = errors.New("mailchimp: API key has not been set")

	// ErrAPIKeyFormat is returned when the provided API key
	// is in an invalid format.
	ErrAPIKeyFormat = errors.New("mailchimp: Invalid API key format")
)

// Error defines a field error.
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// APIError defines the MailChimp API response error structure.
type APIError struct {
	Type   string  `json:"type"`
	Title  string  `json:"title"`
	Status int     `json:"status"`
	Detail string  `json:"detail"`
	Errors []Error `json:"errors,omitempty"`
}

// Error satisfies the error interface method.
func (ae *APIError) Error() string {
	return fmt.Sprintf("mailchimp: API Error: Status: %d Title: %s Detail: %s", ae.Status, ae.Title, ae.Detail)
}
