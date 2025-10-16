package domain

import (
	"testing"
)

func TestEmail_NewEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "user@example.com",
			wantErr: false,
		},
		{
			name:    "valid email with subdomain",
			email:   "user@mail.example.com",
			wantErr: false,
		},
		{
			name:    "valid email with numbers",
			email:   "user123@example123.com",
			wantErr: false,
		},
		{
			name:    "empty email",
			email:   "",
			wantErr: true,
		},
		{
			name:    "invalid email format",
			email:   "invalid-email",
			wantErr: true,
		},
		{
			name:    "email without domain",
			email:   "user@",
			wantErr: true,
		},
		{
			name:    "email without local part",
			email:   "@example.com",
			wantErr: true,
		},
		{
			name:    "email with spaces",
			email:   " user@example.com ",
			wantErr: false, // Should be trimmed and valid
		},
		{
			name:    "email with uppercase",
			email:   "USER@EXAMPLE.COM",
			wantErr: false, // Should be converted to lowercase
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && email != nil {
				// Check that email is normalized (lowercase, trimmed)
				expected := "user@example.com"
				if tt.name == "email with spaces" {
					expected = "user@example.com"
				} else if tt.name == "email with uppercase" {
					expected = "user@example.com"
				} else if tt.name == "valid email with subdomain" {
					expected = "user@mail.example.com"
				} else if tt.name == "valid email with numbers" {
					expected = "user123@example123.com"
				}

				if email.String() != expected {
					t.Errorf("NewEmail() = %v, want %v", email.String(), expected)
				}
			}
		})
	}
}

func TestEmail_Equals(t *testing.T) {
	email1, _ := NewEmail("user@example.com")
	email2, _ := NewEmail("user@example.com")
	email3, _ := NewEmail("different@example.com")

	if !email1.Equals(email2) {
		t.Error("Same emails should be equal")
	}

	if email1.Equals(email3) {
		t.Error("Different emails should not be equal")
	}

	if email1.Equals(nil) {
		t.Error("Email should not equal nil")
	}
}

func TestEmail_String(t *testing.T) {
	email, _ := NewEmail("user@example.com")
	if email.String() != "user@example.com" {
		t.Errorf("String() = %v, want %v", email.String(), "user@example.com")
	}
}
