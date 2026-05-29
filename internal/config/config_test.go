package config

import (
	"os"
	"testing"
)

func TestLoad_Defaults(t *testing.T) {
	// Unset env vars to test defaults
	os.Unsetenv("PORT")
	os.Unsetenv("GIN_MODE")
	os.Unsetenv("MONGODB_URI")
	os.Unsetenv("MONGODB_DATABASE")

	cfg := Load()

	if cfg.Port != "8080" {
		t.Errorf("Port = %q, want %q", cfg.Port, "8080")
	}
	if cfg.GinMode != "release" {
		t.Errorf("GinMode = %q, want %q", cfg.GinMode, "release")
	}
	if cfg.MongoDB.URI != "mongodb://localhost:27017" {
		t.Errorf("MongoDB.URI = %q, want %q", cfg.MongoDB.URI, "mongodb://localhost:27017")
	}
	if cfg.MongoDB.Database != "posts_service" {
		t.Errorf("MongoDB.Database = %q, want %q", cfg.MongoDB.Database, "posts_service")
	}
}

func TestLoad_FromEnv(t *testing.T) {
	tests := []struct {
		name     string
		envKey   string
		envValue string
		check    func(*Config) string
		want     string
	}{
		{
			name:     "custom port",
			envKey:   "PORT",
			envValue: "3000",
			check:    func(c *Config) string { return c.Port },
			want:     "3000",
		},
		{
			name:     "custom gin mode",
			envKey:   "GIN_MODE",
			envValue: "debug",
			check:    func(c *Config) string { return c.GinMode },
			want:     "debug",
		},
		{
			name:     "custom mongodb uri",
			envKey:   "MONGODB_URI",
			envValue: "mongodb://user:pass@host:27017/db",
			check:    func(c *Config) string { return c.MongoDB.URI },
			want:     "mongodb://user:pass@host:27017/db",
		},
		{
			name:     "custom mongodb database",
			envKey:   "MONGODB_DATABASE",
			envValue: "my_custom_db",
			check:    func(c *Config) string { return c.MongoDB.Database },
			want:     "my_custom_db",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.envKey, tt.envValue)
			defer os.Unsetenv(tt.envKey)

			cfg := Load()
			got := tt.check(cfg)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
