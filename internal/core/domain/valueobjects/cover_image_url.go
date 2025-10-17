package valueobjects

import (
	"errors"
	"net/url"
	"strings"
)

// CoverImageURL represents a cover image URL value object
type CoverImageURL struct {
	value string
}

// NewCoverImageURL creates a new CoverImageURL value object with validation
func NewCoverImageURL(imageURL string) (*CoverImageURL, error) {
	imageURL = strings.TrimSpace(imageURL)

	if imageURL == "" {
		return &CoverImageURL{value: ""}, nil // Empty is allowed for optional field
	}

	if !isValidURL(imageURL) {
		return nil, errors.New("invalid URL format")
	}

	if !isValidImageURL(imageURL) {
		return nil, errors.New("URL must point to an image")
	}

	return &CoverImageURL{value: imageURL}, nil
}

// String returns the URL as a string
func (c *CoverImageURL) String() string {
	return c.value
}

// IsEmpty checks if the URL is empty
func (c *CoverImageURL) IsEmpty() bool {
	return c.value == ""
}

// Equals checks if two URLs are equal
func (c *CoverImageURL) Equals(other *CoverImageURL) bool {
	if other == nil {
		return false
	}
	return c.value == other.value
}

// MarshalJSON implements json.Marshaler interface
func (c *CoverImageURL) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.value + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (c *CoverImageURL) UnmarshalJSON(data []byte) error {
	urlStr := strings.Trim(string(data), `"`)

	newURL, err := NewCoverImageURL(urlStr)
	if err != nil {
		return err
	}

	c.value = newURL.value
	return nil
}

// isValidURL checks if the string is a valid URL
func isValidURL(urlStr string) bool {
	_, err := url.Parse(urlStr)
	return err == nil
}

// isValidImageURL checks if the URL points to an image
func isValidImageURL(urlStr string) bool {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Check if URL has a scheme (http or https)
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Check file extension for common image formats
	path := strings.ToLower(parsedURL.Path)

	// Valid image extensions only
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg", ".bmp", ".ico"}

	// Check for valid image extensions
	for _, ext := range imageExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	// Reject anything else (no extension, invalid extensions, etc.)
	return false
}
