package vo

import (
	"testing"
)

func TestTitle_NewTitle(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		wantErr bool
	}{
		{
			name:    "valid title",
			title:   "My Awesome Post",
			wantErr: false,
		},
		{
			name:    "valid short title",
			title:   "Hi",
			wantErr: false,
		},
		{
			name:    "valid long title",
			title:   "This is a very long title that should still be valid because it's under 200 characters and contains meaningful content",
			wantErr: false,
		},
		{
			name:    "empty title",
			title:   "",
			wantErr: true,
		},
		{
			name:    "whitespace only title",
			title:   "   ",
			wantErr: true,
		},
		{
			name:    "title too long",
			title:   "This is an extremely long title that exceeds the maximum allowed length of 200 characters and should fail validation because it's way too long and doesn't provide good user experience and should be rejected by the system",
			wantErr: true,
		},
		{
			name:    "title with leading/trailing spaces",
			title:   "  My Title  ",
			wantErr: false, // Should be trimmed and valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, err := NewTitle(tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && title != nil {
				// Check that title is trimmed
				expected := tt.title
				if tt.name == "title with leading/trailing spaces" {
					expected = "My Title"
				}

				if title.String() != expected {
					t.Errorf("NewTitle() = %v, want %v", title.String(), expected)
				}
			}
		})
	}
}

func TestTitle_Equals(t *testing.T) {
	title1, _ := NewTitle("My Post")
	title2, _ := NewTitle("My Post")
	title3, _ := NewTitle("Different Post")

	if !title1.Equals(title2) {
		t.Error("Same titles should be equal")
	}

	if title1.Equals(title3) {
		t.Error("Different titles should not be equal")
	}

	if title1.Equals(nil) {
		t.Error("Title should not equal nil")
	}
}

func TestTitle_String(t *testing.T) {
	title, _ := NewTitle("My Awesome Post")
	if title.String() != "My Awesome Post" {
		t.Errorf("String() = %v, want %v", title.String(), "My Awesome Post")
	}
}

func TestTitle_Length(t *testing.T) {
	title, _ := NewTitle("Hello World")
	if title.Length() != 11 {
		t.Errorf("Length() = %v, want %v", title.Length(), 11)
	}
}

func TestTitle_IsEmpty(t *testing.T) {
	title, _ := NewTitle("Hello")
	if title.IsEmpty() {
		t.Error("Non-empty title should not be empty")
	}

	emptyTitle, _ := NewTitle("")
	if !emptyTitle.IsEmpty() {
		t.Error("Empty title should be empty")
	}
}
