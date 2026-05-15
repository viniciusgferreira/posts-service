package vo

import (
	"testing"
)

func TestSlug_NewSlugFromTitle(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected string
	}{
		{
			name:     "simple title",
			title:    "My Awesome Post",
			expected: "my-awesome-post",
		},
		{
			name:     "title with special characters",
			title:    "Hello World!",
			expected: "hello-world",
		},
		{
			name:     "title with Portuguese characters",
			title:    "Arquitetura Hexagonal",
			expected: "arquitetura-hexagonal",
		},
		{
			name:     "title with numbers",
			title:    "Post 2024",
			expected: "post-2024",
		},
		{
			name:     "title with multiple spaces",
			title:    "Multiple   Spaces   Here",
			expected: "multiple-spaces-here",
		},
		{
			name:     "title with special characters and spaces",
			title:    "Test @#$%^&*()_+",
			expected: "test",
		},
		{
			name:     "title with leading/trailing spaces",
			title:    "  Trimmed Title  ",
			expected: "trimmed-title",
		},
		{
			name:     "title with consecutive special characters",
			title:    "Title!!!@@@###",
			expected: "title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, _ := NewTitle(tt.title)
			slug := NewSlugFromTitle(title)

			if slug.String() != tt.expected {
				t.Errorf("NewSlugFromTitle() = %v, want %v", slug.String(), tt.expected)
			}
		})
	}
}

func TestSlug_NewSlug(t *testing.T) {
	slug := NewSlug("my-slug")
	if slug.String() != "my-slug" {
		t.Errorf("NewSlug() = %v, want %v", slug.String(), "my-slug")
	}
}

func TestSlug_Equals(t *testing.T) {
	slug1 := NewSlug("my-slug")
	slug2 := NewSlug("my-slug")
	slug3 := NewSlug("different-slug")

	if !slug1.Equals(slug2) {
		t.Error("Same slugs should be equal")
	}

	if slug1.Equals(slug3) {
		t.Error("Different slugs should not be equal")
	}

	if slug1.Equals(nil) {
		t.Error("Slug should not equal nil")
	}
}

func TestSlug_String(t *testing.T) {
	slug := NewSlug("my-awesome-slug")
	if slug.String() != "my-awesome-slug" {
		t.Errorf("String() = %v, want %v", slug.String(), "my-awesome-slug")
	}
}

func TestSlug_IsEmpty(t *testing.T) {
	slug := NewSlug("not-empty")
	if slug.IsEmpty() {
		t.Error("Non-empty slug should not be empty")
	}

	emptySlug := NewSlug("")
	if !emptySlug.IsEmpty() {
		t.Error("Empty slug should be empty")
	}
}

func TestSlug_Value(t *testing.T) {
	slug := NewSlug("test-slug")
	if slug.String() != "test-slug" {
		t.Errorf("String() = %v, want %v", slug.String(), "test-slug")
	}
}
