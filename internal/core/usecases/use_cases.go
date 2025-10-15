package usecases

import (
	"fmt"
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/ports"
)

type PostUseCases struct {
	postRepo   ports.PostRepository
	authorRepo ports.AuthorRepository
	notifier   ports.NotificationService
}

func NewPostUseCases(postRepo ports.PostRepository, authorRepo ports.AuthorRepository, notifier ports.NotificationService) *PostUseCases {
	return &PostUseCases{
		postRepo:   postRepo,
		authorRepo: authorRepo,
		notifier:   notifier,
	}
}

func (uc *PostUseCases) CreatePost(title, content, authorID, coverImageURL string) (*domain.Post, error) {
	author, err := uc.authorRepo.FindByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	postID := uc.generatePostID()
	post, err := domain.NewPost(postID, title, content, author, coverImageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	uniqueSlug, err := uc.ensureSlugUniqueness(post.Slug)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure slug uniqueness: %w", err)
	}
	post.Slug = uniqueSlug

	if err := uc.postRepo.Save(post); err != nil {
		return nil, fmt.Errorf("failed to save post: %w", err)
	}

	// Business rule: Send notification (mocked)
	if err := uc.notifier.NotifyPostCreated(post); err != nil {
		// Log error but don't fail the operation
		// TODO: Implement proper logging
		fmt.Printf("Warning: Failed to send notification for post %s: %v\n", post.ID, err)
	}

	return post, nil
}

func (uc *PostUseCases) GetPost(id string) (*domain.Post, error) {
	// TODO: Implement get post logic
	return nil, fmt.Errorf("get post not implemented")
}

func (uc *PostUseCases) GetAllPosts() ([]*domain.Post, error) {
	// TODO: Implement get all posts logic
	return nil, fmt.Errorf("get all posts not implemented")
}

func (uc *PostUseCases) UpdatePost(id, title, content, coverImageURL string) (*domain.Post, error) {
	// TODO: Implement update post logic
	return nil, fmt.Errorf("update post not implemented")
}

func (uc *PostUseCases) DeletePost(id string) error {
	// TODO: Implement delete post logic
	return fmt.Errorf("delete post not implemented")
}

func (uc *PostUseCases) ensureSlugUniqueness(baseSlug string) (string, error) {
	slug := baseSlug
	counter := 1

	for {
		existingPost, err := uc.postRepo.FindBySlug(slug)
		if err != nil {
			return slug, nil
		}

		if existingPost == nil {
			return slug, nil
		}

		// Slug already exists, try with numeric suffix
		counter++
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)

		// Avoid infinite loop
		if counter > 1000 {
			return "", fmt.Errorf("unable to generate unique slug after 1000 attempts")
		}
	}
}

// generatePostID generates a unique ID for the post
func (uc *PostUseCases) generatePostID() string {
	// Use timestamp to generate unique ID
	return fmt.Sprintf("post_%d", time.Now().UnixNano())
}
