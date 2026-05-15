package vo

import (
	"strings"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
)

// MarkdownContent represents markdown content value object
type MarkdownContent struct {
	value string
}

// NewMarkdownContent creates a new MarkdownContent value object with validation
func NewMarkdownContent(content string) (*MarkdownContent, error) {
	content = strings.TrimSpace(content)

	if content == "" {
		return nil, errs.ContentEmpty
	}

	if len(content) < 10 {
		return nil, errs.ContentTooShort
	}

	if len(content) > 50000 {
		return nil, errs.ContentTooLong
	}

	return &MarkdownContent{value: content}, nil
}

// String returns the content as a string
func (m *MarkdownContent) String() string {
	return m.value
}

// Length returns the length of the content
func (m *MarkdownContent) Length() int {
	return len(m.value)
}

// Equals checks if two markdown contents are equal
func (m *MarkdownContent) Equals(other *MarkdownContent) bool {
	if other == nil {
		return false
	}
	return m.value == other.value
}

// IsEmpty checks if the content is empty
func (m *MarkdownContent) IsEmpty() bool {
	return m.value == ""
}

// HasMinimumLength checks if content meets minimum length requirement
func (m *MarkdownContent) HasMinimumLength() bool {
	return len(m.value) >= 10
}
