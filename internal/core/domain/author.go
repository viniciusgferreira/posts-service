package domain

import (
	"errors"
	"strings"
)

type Author struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email *Email `json:"email"`
}

func NewAuthor(id, name, email string) (*Author, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("author name cannot be empty")
	}

	validatedEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &Author{
		ID:    id,
		Name:  strings.TrimSpace(name),
		Email: validatedEmail,
	}, nil
}
