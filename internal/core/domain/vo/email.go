package vo

import (
	"regexp"
	"strings"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
)

// Email represents an email address value object
type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return nil, errs.EmailEmpty
	}

	email = strings.ToLower(email)

	if !isValidEmail(email) {
		return nil, errs.EmailInvalid
	}

	return &Email{value: email}, nil
}

func (e *Email) String() string {
	return e.value
}

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
