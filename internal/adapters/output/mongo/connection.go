package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connectTimeout = 10 * time.Second

// NewConnection creates a new MongoDB connection and returns the client and database
func NewConnection(uri, database string, logger *logrus.Logger) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connect: %w", err)
	}

	// Verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, fmt.Errorf("mongo ping: %w", err)
	}

	logger.WithFields(logrus.Fields{
		"database": database,
	}).Info("Connected to MongoDB")

	db := client.Database(database)
	return client, db, nil
}

// Disconnect gracefully closes the MongoDB connection
func Disconnect(client *mongo.Client, logger *logrus.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		logger.WithError(err).Error("Failed to disconnect from MongoDB")
		return
	}
	logger.Info("Disconnected from MongoDB")
}
