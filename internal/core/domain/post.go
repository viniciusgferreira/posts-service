package domain

import (
	"errors"
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/valueobjects"
)

type Post struct {
	ID              string                        `json:"id"`
	Title           *valueobjects.Title           `json:"title"`
	Slug            *valueobjects.Slug            `json:"slug"`
	Author          *Author                       `json:"author"`
	CoverImageURL   *valueobjects.CoverImageURL   `json:"cover_image_url"`
	MarkdownContent *valueobjects.MarkdownContent `json:"markdown_content"`
	CreatedAt       time.Time                     `json:"created_at"`
	UpdatedAt       time.Time                     `json:"updated_at"`
}

func NewPost(id, title, markdownContent string, author *Author, coverImageURL string) (*Post, error) {
	validatedTitle, err := valueobjects.NewTitle(title)
	if err != nil {
		return nil, err
	}

	validatedContent, err := valueobjects.NewMarkdownContent(markdownContent)
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, errors.New("post must have an author")
	}

	slug := valueobjects.NewSlugFromTitle(validatedTitle)

	validatedCoverURL, err := valueobjects.NewCoverImageURL(coverImageURL)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	post := &Post{
		ID:              id,
		Title:           validatedTitle,
		Slug:            slug,
		Author:          author,
		CoverImageURL:   validatedCoverURL,
		MarkdownContent: validatedContent,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	return post, nil
}
