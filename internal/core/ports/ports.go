package ports

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

type PostPort interface {
	Save(post *domain.Post) error
	FindByID(id string) (*domain.Post, error)
	FindAll() ([]*domain.Post, error)
	Update(post *domain.Post) error
	Delete(id string) error
}

type AuthorPort interface {
	Save(author *domain.Author) error
	FindByID(id string) (*domain.Author, error)
	FindAll() ([]*domain.Author, error)
	Update(author *domain.Author) error
	Delete(id string) error
}