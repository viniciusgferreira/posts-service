package ports

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

type PostCreationPort interface {
	Save(post *domain.Post) error
}
