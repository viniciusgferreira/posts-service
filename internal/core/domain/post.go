package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/valueobjects"
)

type Post struct {
	ID              string              `json:"id"`
	Title           *valueobjects.Title `json:"title"`
	Slug            *valueobjects.Slug  `json:"slug"`
	Author          *Author             `json:"author"`
	CoverImageURL   string              `json:"cover_image_url"`
	MarkdownContent string              `json:"markdown_content"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

func NewPost(id, title, markdownContent string, author *Author, coverImageURL string) (*Post, error) {
	validatedTitle, err := valueobjects.NewTitle(title)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(markdownContent) == "" {
		return nil, errors.New("post content cannot be empty")
	}

	if author == nil {
		return nil, errors.New("post must have an author")
	}

	slug := valueobjects.NewSlugFromTitle(validatedTitle)

	now := time.Now()
	post := &Post{
		ID:              id,
		Title:           validatedTitle,
		Slug:            slug,
		Author:          author,
		CoverImageURL:   strings.TrimSpace(coverImageURL),
		MarkdownContent: strings.TrimSpace(markdownContent),
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	return post, nil
}
