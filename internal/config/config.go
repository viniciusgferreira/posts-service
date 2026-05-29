package config

import "os"

// Config holds all application configuration
type Config struct {
	Port    string
	GinMode string
	MongoDB MongoDBConfig
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "release"),
		MongoDB: loadMongoDBConfig(),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
