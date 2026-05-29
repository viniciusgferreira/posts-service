package mongo

import (
	"context"
	"fmt"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const postsCollection = "posts"

// PostRepository implements ports.PostCreationPort using MongoDB
type PostRepository struct {
	collection *mongo.Collection
}

// NewPostRepository creates a new PostRepository
func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{
		collection: db.Collection(postsCollection),
	}
}

// Save persists a Post to MongoDB and sets the generated ID on the entity
func (r *PostRepository) Save(post *domain.Post) error {
	doc := toDocument(post)

	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errs.DuplicateError
		}
		return fmt.Errorf("mongo insert post: %w", err)
	}

	// Set the generated MongoDB ObjectID back on the domain entity
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		post.ID = oid.Hex()
	}

	return nil
}
