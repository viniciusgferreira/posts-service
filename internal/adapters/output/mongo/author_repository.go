package mongo

import (
	"context"
	"fmt"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const authorsCollection = "authors"

// AuthorRepository implements ports.AuthorReadingPort using MongoDB
type AuthorRepository struct {
	collection *mongo.Collection
}

// NewAuthorRepository creates a new AuthorRepository
func NewAuthorRepository(db *mongo.Database) *AuthorRepository {
	return &AuthorRepository{
		collection: db.Collection(authorsCollection),
	}
}

// FindByID retrieves an Author by their ID from MongoDB
func (r *AuthorRepository) FindByID(id string) (*domain.Author, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid author id: %w", errs.AuthorNotFound)
	}

	var doc authorDocument
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.AuthorNotFound
		}
		return nil, fmt.Errorf("mongo find author: %w", err)
	}

	return toAuthorDomain(doc)
}
