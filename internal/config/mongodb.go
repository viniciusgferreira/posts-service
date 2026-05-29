package config

// MongoDBConfig holds MongoDB connection configuration
type MongoDBConfig struct {
	URI      string
	Database string
}

// loadMongoDBConfig reads MongoDB configuration from environment variables
func loadMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		Database: getEnv("MONGODB_DATABASE", "posts_service"),
	}
}
