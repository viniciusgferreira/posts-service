package mock

import (
	"fmt"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
)

// MockPostRepository implements PostRepository interface for testing/development
type MockPostRepository struct {
	posts  map[string]*domain.Post
	slugs  map[string]*domain.Post
	titles map[string]*domain.Post // Track titles for uniqueness check
}

func NewMockPostRepository() *MockPostRepository {
	return &MockPostRepository{
		posts:  make(map[string]*domain.Post),
		slugs:  make(map[string]*domain.Post),
		titles: make(map[string]*domain.Post),
	}
}

func (r *MockPostRepository) Save(post *domain.Post) error {
	// Check for duplicate title (case-insensitive)
	titleKey := post.Title.String()
	for existingTitle, existingPost := range r.titles {
		// Skip if it's the same post being updated
		if existingPost.ID == post.ID {
			continue
		}
		// Check for exact title match (case-insensitive)
		if existingTitle == titleKey {
			return errs.DuplicateError
		}
	}

	// Check for duplicate slug
	slugKey := post.Slug.String()
	if existingPost, exists := r.slugs[slugKey]; exists && existingPost.ID != post.ID {
		return errs.DuplicateError
	}

	// Save the post
	r.posts[post.ID] = post
	r.slugs[slugKey] = post
	r.titles[titleKey] = post
	return nil
}

func (r *MockPostRepository) FindByID(id string) (*domain.Post, error) {
	post, exists := r.posts[id]
	if !exists {
		return nil, fmt.Errorf("post with id %s not found", id)
	}
	return post, nil
}

func (r *MockPostRepository) FindBySlug(slug string) (*domain.Post, error) {
	post, exists := r.slugs[slug]
	if !exists {
		return nil, fmt.Errorf("post with slug %s not found", slug)
	}
	return post, nil
}

func (r *MockPostRepository) FindAll() ([]*domain.Post, error) {
	posts := make([]*domain.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *MockPostRepository) Update(post *domain.Post) error {
	// Get existing post to clean up old title/slug mappings
	existingPost, exists := r.posts[post.ID]
	if exists {
		// Remove old title and slug mappings
		delete(r.titles, existingPost.Title.String())
		delete(r.slugs, existingPost.Slug.String())
	}

	// Check for duplicate title (case-insensitive)
	titleKey := post.Title.String()
	for existingTitle, existingPost := range r.titles {
		// Skip if it's the same post being updated
		if existingPost.ID == post.ID {
			continue
		}
		// Check for exact title match
		if existingTitle == titleKey {
			return errs.DuplicateError
		}
	}

	// Check for duplicate slug
	slugKey := post.Slug.String()
	if existingPost, exists := r.slugs[slugKey]; exists && existingPost.ID != post.ID {
		return errs.DuplicateError
	}

	// Update the post
	r.posts[post.ID] = post
	r.slugs[slugKey] = post
	r.titles[titleKey] = post
	return nil
}

func (r *MockPostRepository) Delete(id string) error {
	post, exists := r.posts[id]
	if !exists {
		return fmt.Errorf("post with id %s not found", id)
	}
	delete(r.posts, id)
	delete(r.slugs, post.Slug.String())
	delete(r.titles, post.Title.String())
	return nil
}

// MockAuthorRepository implements AuthorRepository interface for testing/development
type MockAuthorRepository struct {
	authors map[string]*domain.Author
}

func NewMockAuthorRepository() *MockAuthorRepository {
	return &MockAuthorRepository{
		authors: make(map[string]*domain.Author),
	}
}

func (r *MockAuthorRepository) Save(author *domain.Author) error {
	r.authors[author.ID] = author
	return nil
}

func (r *MockAuthorRepository) FindByID(id string) (*domain.Author, error) {
	author, exists := r.authors[id]
	if !exists {
		return nil, fmt.Errorf("author with id %s not found", id)
	}
	return author, nil
}

func (r *MockAuthorRepository) FindAll() ([]*domain.Author, error) {
	authors := make([]*domain.Author, 0, len(r.authors))
	for _, author := range r.authors {
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *MockAuthorRepository) Update(author *domain.Author) error {
	r.authors[author.ID] = author
	return nil
}

func (r *MockAuthorRepository) Delete(id string) error {
	_, exists := r.authors[id]
	if !exists {
		return fmt.Errorf("author with id %s not found", id)
	}
	delete(r.authors, id)
	return nil
}

// MockNotificationService implements NotificationService interface for testing/development
type MockNotificationService struct{}

func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{}
}

func (s *MockNotificationService) NotifyPostCreated(post *domain.Post) error {
	fmt.Printf("🔔 NOTIFICATION: New post created!\n")
	fmt.Printf("   Title: %s\n", post.Title)
	fmt.Printf("   Author: %s (%s)\n", post.Author.Name, post.Author.Email.String())
	fmt.Printf("   Slug: %s\n", post.Slug.String())
	fmt.Printf("   ID: %s\n", post.ID)
	fmt.Printf("   Created at: %s\n", post.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("=====================================\n")
	return nil
}
