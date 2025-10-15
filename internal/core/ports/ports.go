package ports

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

type PostRepository interface {
	Save(post *domain.Post) error
	FindByID(id string) (*domain.Post, error)
	FindAll() ([]*domain.Post, error)
	Update(post *domain.Post) error
	Delete(id string) error
}

type AuthorRepository interface {
	Save(author *domain.Author) error
	FindByID(id string) (*domain.Author, error)
	FindAll() ([]*domain.Author, error)
	Update(author *domain.Author) error
	Delete(id string) error
}

// TODO: This will be implemented in the use cases layer
type PostUseCases interface {
	CreatePost(title, content, authorID, coverImageURL string) (*domain.Post, error)
	GetPost(id string) (*domain.Post, error)
	GetAllPosts() ([]*domain.Post, error)
	UpdatePost(id, title, content, coverImageURL string) (*domain.Post, error)
	DeletePost(id string) error
}
