package domain

import (
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
)

type Post struct {
	ID              string              `json:"id"`
	Title           *vo.Title           `json:"title"`
	Slug            *vo.Slug            `json:"slug"`
	Author          *Author             `json:"author"`
	CoverImageURL   *vo.CoverImageURL   `json:"cover_image_url"`
	MarkdownContent *vo.MarkdownContent `json:"markdown_content"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

func NewPost(id, title, markdownContent string, author *Author, coverImageURL string) (*Post, error) {
	validatedTitle, err := vo.NewTitle(title)
	if err != nil {
		return nil, err
	}

	validatedContent, err := vo.NewMarkdownContent(markdownContent)
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, errs.MissingAuthor
	}

	slug := vo.NewSlugFromTitle(validatedTitle)

	validatedCoverURL, err := vo.NewCoverImageURL(coverImageURL)
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
