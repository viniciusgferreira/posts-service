package valueobjects

import (
	"testing"
)

func TestCoverImageURL_NewCoverImageURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid image URL with jpg extension",
			url:     "https://example.com/image.jpg",
			wantErr: false,
		},
		{
			name:    "valid image URL with png extension",
			url:     "https://example.com/image.png",
			wantErr: false,
		},
		{
			name:    "valid image URL with gif extension",
			url:     "https://example.com/image.gif",
			wantErr: false,
		},
		{
			name:    "valid image URL with webp extension",
			url:     "https://example.com/image.webp",
			wantErr: false,
		},
		{
			name:    "valid image URL with svg extension",
			url:     "https://example.com/image.svg",
			wantErr: false,
		},
		{
			name:    "valid image URL without extension",
			url:     "https://example.com/api/image/123",
			wantErr: true, // Now rejects URLs without image extensions
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: false, // Empty is allowed
		},
		{
			name:    "URL with spaces",
			url:     "  https://example.com/image.jpg  ",
			wantErr: false, // Should be trimmed
		},
		{
			name:    "invalid URL format",
			url:     "not-a-url",
			wantErr: true,
		},
		{
			name:    "URL without scheme",
			url:     "example.com/image.jpg",
			wantErr: true,
		},
		{
			name:    "URL with ftp scheme",
			url:     "ftp://example.com/image.jpg",
			wantErr: true,
		},
		{
			name:    "URL with PDF extension",
			url:     "https://example.com/document.pdf",
			wantErr: true,
		},
		{
			name:    "URL with DOC extension",
			url:     "https://example.com/document.doc",
			wantErr: true,
		},
		{
			name:    "URL with TXT extension",
			url:     "https://example.com/document.txt",
			wantErr: true,
		},
		{
			name:    "URL with ZIP extension",
			url:     "https://example.com/archive.zip",
			wantErr: true,
		},
		{
			name:    "URL with MP4 extension",
			url:     "https://example.com/video.mp4",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coverURL, err := NewCoverImageURL(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCoverImageURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && coverURL != nil {
				expected := tt.url
				if tt.name == "URL with spaces" {
					expected = "https://example.com/image.jpg"
				}

				if coverURL.String() != expected {
					t.Errorf("NewCoverImageURL() = %v, want %v", coverURL.String(), expected)
				}
			}
		})
	}
}

func TestCoverImageURL_Equals(t *testing.T) {
	url1, _ := NewCoverImageURL("https://example.com/image.jpg")
	url2, _ := NewCoverImageURL("https://example.com/image.jpg")
	url3, _ := NewCoverImageURL("https://example.com/different.jpg")

	if !url1.Equals(url2) {
		t.Error("Same URLs should be equal")
	}

	if url1.Equals(url3) {
		t.Error("Different URLs should not be equal")
	}

	if url1.Equals(nil) {
		t.Error("URL should not equal nil")
	}
}

func TestCoverImageURL_String(t *testing.T) {
	coverURL, _ := NewCoverImageURL("https://example.com/image.jpg")
	if coverURL.String() != "https://example.com/image.jpg" {
		t.Errorf("String() = %v, want %v", coverURL.String(), "https://example.com/image.jpg")
	}
}

func TestCoverImageURL_IsEmpty(t *testing.T) {
	coverURL, _ := NewCoverImageURL("https://example.com/image.jpg")
	if coverURL.IsEmpty() {
		t.Error("Non-empty URL should not be empty")
	}

	emptyURL, _ := NewCoverImageURL("")
	if !emptyURL.IsEmpty() {
		t.Error("Empty URL should be empty")
	}
}

func TestCoverImageURL_EmptyURL(t *testing.T) {
	emptyURL, err := NewCoverImageURL("")
	if err != nil {
		t.Errorf("Empty URL should not return error, got %v", err)
	}
	if !emptyURL.IsEmpty() {
		t.Error("Empty URL should be empty")
	}
}
