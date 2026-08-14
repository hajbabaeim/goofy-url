package domain

import (
	"errors"
	"net/url"
	"time"
)

type ShortURL struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Original  string    `json:"original" validate:"required,url"`
	Expires   time.Time `json:"expires"`
}

func ValidateURL(raw string) error {
	if _, err := url.ParseRequestURI(raw); err != nil {
		return errors.New("invalid URL format")
	}
	return nil
}
