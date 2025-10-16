package domain

import (
	"errors"
	"strings"
	"time"
)

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     *Email    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAuthor(id, name, email string) (*Author, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("author name cannot be empty")
	}

	validatedEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Author{
		ID:        id,
		Name:      strings.TrimSpace(name),
		Email:     validatedEmail,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
