package ports

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

type AuthorReadingPort interface {
	FindByID(id string) (*domain.Author, error)
}
