package mongo

import (
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// postDocument represents a post as stored in MongoDB
type postDocument struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title"`
	Slug            string             `bson:"slug"`
	AuthorID        string             `bson:"author_id"`
	AuthorName      string             `bson:"author_name"`
	AuthorEmail     string             `bson:"author_email"`
	CoverImageURL   string             `bson:"cover_image_url"`
	MarkdownContent string             `bson:"markdown_content"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

// toDocument converts a domain Post to a MongoDB document
func toDocument(post *domain.Post) postDocument {
	doc := postDocument{
		Title:           post.Title.String(),
		Slug:            post.Slug.String(),
		AuthorID:        post.Author.ID,
		AuthorName:      post.Author.Name,
		AuthorEmail:     post.Author.Email.String(),
		CoverImageURL:   post.CoverImageURL.String(),
		MarkdownContent: post.MarkdownContent.String(),
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
	}

	// If the post already has an ID, use it
	if post.ID != "" {
		objID, err := primitive.ObjectIDFromHex(post.ID)
		if err == nil {
			doc.ID = objID
		}
	}

	return doc
}

// toDomain converts a MongoDB document to a domain Post
func toDomain(doc postDocument) (*domain.Post, error) {
	email, err := vo.NewEmail(doc.AuthorEmail)
	if err != nil {
		return nil, err
	}

	author := &domain.Author{
		ID:    doc.AuthorID,
		Name:  doc.AuthorName,
		Email: email,
	}

	post, err := domain.NewPostBuilder().
		WithID(doc.ID.Hex()).
		WithTitle(doc.Title).
		WithSlug(doc.Slug).
		WithMarkdownContent(doc.MarkdownContent).
		WithAuthor(author).
		WithCoverImageURL(doc.CoverImageURL).
		WithCreatedAt(doc.CreatedAt).
		WithUpdatedAt(doc.UpdatedAt).
		Build()

	if err != nil {
		return nil, err
	}

	return post, nil
}

// authorDocument represents an author as stored in MongoDB
type authorDocument struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

// toAuthorDomain converts a MongoDB author document to a domain Author
func toAuthorDomain(doc authorDocument) (*domain.Author, error) {
	email, err := vo.NewEmail(doc.Email)
	if err != nil {
		return nil, err
	}

	return &domain.Author{
		ID:    doc.ID.Hex(),
		Name:  doc.Name,
		Email: email,
	}, nil
}
