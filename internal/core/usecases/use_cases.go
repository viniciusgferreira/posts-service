package usecases

import (
	"errors"
	"fmt"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"github.com/viniciusgferreira/posts-service/internal/core/ports"
)

// PostUseCases handles all post-related business logic
type PostUseCases struct {
	postCreationPort  ports.PostCreationPort
	authorReadingPort ports.AuthorReadingPort
}

// NewPostUseCases creates a new instance of PostUseCases
func NewPostUseCases(postRepository ports.PostCreationPort, authorRepository ports.AuthorReadingPort) *PostUseCases {
	return &PostUseCases{
		postCreationPort:  postRepository,
		authorReadingPort: authorRepository,
	}
}

// CreatePost creates a new post from a domain Post entity
func (uc *PostUseCases) CreatePost(post *domain.Post) (*domain.Post, error) {
	if err := uc.postCreationPort.Save(post); err != nil {
		return nil, fmt.Errorf("failed to save post: %w", err)
	}

	return post, nil
}

// GetAuthor retrieves an author by ID
func (uc *PostUseCases) GetAuthor(authorID string) (*domain.Author, error) {
	author, err := uc.authorReadingPort.FindByID(authorID)
	if err != nil {
		var appErr errs.AppErrorInterface
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	return author, nil
}
