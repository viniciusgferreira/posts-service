package vo

import (
	"net/url"
	"strings"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
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
		return nil, errs.CoverURLInvalid
	}

	if !isValidImageURL(imageURL) {
		return nil, errs.CoverURLNotImage
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
