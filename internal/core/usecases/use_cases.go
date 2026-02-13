package usecases

import (
	"fmt"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"github.com/viniciusgferreira/posts-service/internal/core/ports"
)

// PostUseCases handles all post-related business logic
type PostUseCases struct {
	postCreationPort ports.PostCreationPort
	authorReadingPort ports.AuthorReadingPort
}

// NewPostUseCases creates a new instance of PostUseCases
func NewPostUseCases(postRepository ports.PostCreationPort, authorRepository ports.AuthorReadingPort) *PostUseCases {
	return &PostUseCases{
		postCreationPort: postRepository,
		authorReadingPort: authorRepository,
	}
}

// CreatePost creates a new post
func (uc *PostUseCases) CreatePost(title, markdownContent, authorID, coverImageURL string) (*domain.Post, error) {
	// Get author by ID
	author, err := uc.authorReadingPort.FindByID(authorID)
	if err != nil {
		return nil, errs.AuthorNotFound
	}

	// Create post domain entity (this will validate all fields)
	post, err := domain.NewPost(title, markdownContent, author, coverImageURL)
	if err != nil {
		return nil, err
	}

	// Save post to repository
	if err := uc.postCreationPort.Save(post); err != nil {
		return nil, fmt.Errorf("failed to save post: %w", err)
	}

	return post, nil
}
