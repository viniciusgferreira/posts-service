package ports

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

type AuthorCreationPort interface {
	FindByID(id string) (*domain.Author, error)
}
