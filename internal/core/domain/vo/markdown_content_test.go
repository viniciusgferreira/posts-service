package vo

import (
	"strings"
	"testing"
)

func TestMarkdownContent_NewMarkdownContent(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "valid markdown content",
			content: "# My Post\n\nThis is a valid markdown content with proper length.",
			wantErr: false,
		},
		{
			name:    "valid short content",
			content: "Short content",
			wantErr: false,
		},
		{
			name:    "valid long content",
			content: strings.Repeat("This is a long markdown content. ", 100),
			wantErr: false,
		},
		{
			name:    "empty content",
			content: "",
			wantErr: true,
		},
		{
			name:    "whitespace only content",
			content: "   ",
			wantErr: true,
		},
		{
			name:    "content too short",
			content: "Short",
			wantErr: true,
		},
		{
			name:    "content with leading/trailing spaces",
			content: "  Valid content here  ",
			wantErr: false, // Should be trimmed and valid
		},
		{
			name:    "content with markdown syntax",
			content: "# Title\n\n**Bold text** and *italic text*\n\n- List item 1\n- List item 2",
			wantErr: false,
		},
		{
			name:    "content at minimum length",
			content: "1234567890", // Exactly 10 characters
			wantErr: false,
		},
		{
			name:    "content just under minimum",
			content: "123456789", // 9 characters
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := NewMarkdownContent(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMarkdownContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && content != nil {
				// Just verify that content was created successfully
				if content.String() == "" {
					t.Error("Content should not be empty")
				}
			}
		})
	}
}

func TestMarkdownContent_Equals(t *testing.T) {
	content1, _ := NewMarkdownContent("This is markdown content")
	content2, _ := NewMarkdownContent("This is markdown content")
	content3, _ := NewMarkdownContent("Different markdown content")

	if !content1.Equals(content2) {
		t.Error("Same contents should be equal")
	}

	if content1.Equals(content3) {
		t.Error("Different contents should not be equal")
	}

	if content1.Equals(nil) {
		t.Error("Content should not equal nil")
	}
}

func TestMarkdownContent_String(t *testing.T) {
	content, _ := NewMarkdownContent("This is markdown content")
	if content.String() != "This is markdown content" {
		t.Errorf("String() = %v, want %v", content.String(), "This is markdown content")
	}
}

func TestMarkdownContent_Length(t *testing.T) {
	content, _ := NewMarkdownContent("Hello World")
	if content.Length() != 11 {
		t.Errorf("Length() = %v, want %v", content.Length(), 11)
	}
}

func TestMarkdownContent_IsEmpty(t *testing.T) {
	content, _ := NewMarkdownContent("Non-empty content")
	if content.IsEmpty() {
		t.Error("Non-empty content should not be empty")
	}
}

func TestMarkdownContent_HasMinimumLength(t *testing.T) {
	content, _ := NewMarkdownContent("This is long enough")
	if !content.HasMinimumLength() {
		t.Error("Content should have minimum length")
	}

	shortContent, _ := NewMarkdownContent("1234567890") // Exactly 10 chars
	if !shortContent.HasMinimumLength() {
		t.Error("Content with exactly 10 characters should have minimum length")
	}
}

func TestMarkdownContent_MaximumLength(t *testing.T) {
	// Test content that's too long (over 50000 characters)
	longContent := strings.Repeat("This is a very long content. ", 2000) // ~50,000+ chars

	_, err := NewMarkdownContent(longContent)
	if err == nil {
		t.Error("Content that's too long should return error")
	}

	if !strings.Contains(err.Error(), "cannot exceed 50000 characters") {
		t.Errorf("Expected error about maximum length, got: %v", err)
	}
}
