package domain

import (
	"errors"
	"strings"
)

type Title struct {
	value string
}

// NewTitle creates a new Title value object with validation
func NewTitle(title string) (*Title, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	title = strings.TrimSpace(title)

	if len(title) < 1 {
		return nil, errors.New("title cannot be empty")
	}

	if len(title) > 200 {
		return nil, errors.New("title cannot exceed 200 characters")
	}

	return &Title{value: title}, nil
}

// String returns the title as a string
func (t *Title) String() string {
	return t.value
}

// Length returns the length of the title
func (t *Title) Length() int {
	return len(t.value)
}

// Equals checks if two titles are equal
func (t *Title) Equals(other *Title) bool {
	if other == nil {
		return false
	}
	return t.value == other.value
}

// IsEmpty checks if the title is empty
func (t *Title) IsEmpty() bool {
	return t.value == ""
}

// MarshalJSON implements json.Marshaler interface
func (t *Title) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.value + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Title) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	titleStr := strings.Trim(string(data), `"`)

	newTitle, err := NewTitle(titleStr)
	if err != nil {
		return err
	}

	t.value = newTitle.value
	return nil
}
