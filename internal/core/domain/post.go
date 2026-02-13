package domain

import (
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
)

type Post struct {
	ID              string
	Title           *vo.Title
	Slug            *vo.Slug
	Author          *Author
	CoverImageURL   *vo.CoverImageURL
	MarkdownContent *vo.MarkdownContent
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewPost(title, markdownContent string, author *Author, coverImageURL string) (*Post, error) {
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
