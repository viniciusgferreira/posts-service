package domain

import (
	"errors"
	"regexp"
	"strings"
)

// Email represents an email address value object
type Email struct {
	value string
}

// NewEmail creates a new Email value object with validation
func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	return &Email{value: email}, nil
}

// String returns the email as a string
func (e *Email) String() string {
	return e.value
}

// Value returns the email value (alias for String)
func (e *Email) Value() string {
	return e.value
}

// Equals checks if two emails are equal
func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func isValidEmail(email string) bool {
	// RFC 5322 compliant email regex (simplified but robust)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (e *Email) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.value + `"`), nil
}

func (e *Email) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	emailStr := strings.Trim(string(data), `"`)

	newEmail, err := NewEmail(emailStr)
	if err != nil {
		return err
	}

	e.value = newEmail.value
	return nil
}
