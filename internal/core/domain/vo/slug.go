package vo

import (
	"strings"
)

// Slug represents a URL-friendly slug value object
type Slug struct {
	value string
}

// NewSlugFromTitle creates a new Slug value object from a Title
func NewSlugFromTitle(title *Title) *Slug {
	slug := generateSlugFromTitle(title.String())
	return &Slug{value: slug}
}

// NewSlug creates a new Slug value object from a string
func NewSlug(slug string) *Slug {
	return &Slug{value: slug}
}

// String returns the slug as a string
func (s *Slug) String() string {
	return s.value
}

// Equals checks if two slugs are equal
func (s *Slug) Equals(other *Slug) bool {
	if other == nil {
		return false
	}
	return s.value == other.value
}

// IsEmpty checks if the slug is empty
func (s *Slug) IsEmpty() bool {
	return s.value == ""
}

// MarshalJSON implements json.Marshaler interface
func (s *Slug) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.value + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (s *Slug) UnmarshalJSON(data []byte) error {
	slugStr := strings.Trim(string(data), `"`)
	s.value = slugStr
	return nil
}

// generateSlugFromTitle creates a URL-friendly slug from a title
func generateSlugFromTitle(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Replace Portuguese characters
	slug = strings.ReplaceAll(slug, "ç", "c")
	slug = strings.ReplaceAll(slug, "ã", "a")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "à", "a")
	slug = strings.ReplaceAll(slug, "â", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "ê", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ô", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ü", "u")

	// Keep only alphanumeric characters and hyphens
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	slug = result.String()

	// Remove multiple consecutive hyphens
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}
