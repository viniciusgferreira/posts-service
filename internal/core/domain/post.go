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

type postBuilder struct {
	id              string
	title           string
	slug            string
	markdownContent string
	author          *Author
	coverImageURL   string
	createdAt       time.Time
	updatedAt       time.Time
}

// NewPostBuilder creates a new Post builder
func NewPostBuilder() *postBuilder {
	return &postBuilder{}
}

// WithID sets the post ID
func (b *postBuilder) WithID(id string) *postBuilder {
	b.id = id
	return b
}

// WithTitle sets the post title
func (b *postBuilder) WithTitle(title string) *postBuilder {
	b.title = title
	return b
}

// WithSlug sets the post slug (optional; if empty, derived from title)
func (b *postBuilder) WithSlug(slug string) *postBuilder {
	b.slug = slug
	return b
}

// WithMarkdownContent sets the markdown content
func (b *postBuilder) WithMarkdownContent(content string) *postBuilder {
	b.markdownContent = content
	return b
}

// WithAuthor sets the post author
func (b *postBuilder) WithAuthor(author *Author) *postBuilder {
	b.author = author
	return b
}

// WithCoverImageURL sets the cover image URL
func (b *postBuilder) WithCoverImageURL(url string) *postBuilder {
	b.coverImageURL = url
	return b
}

// WithCreatedAt sets the creation timestamp
func (b *postBuilder) WithCreatedAt(t time.Time) *postBuilder {
	b.createdAt = t
	return b
}

// WithUpdatedAt sets the update timestamp
func (b *postBuilder) WithUpdatedAt(t time.Time) *postBuilder {
	b.updatedAt = t
	return b
}

// Build validates and constructs the Post entity
func (b *postBuilder) Build() (*Post, error) {
	// Validate and create title value object
	validatedTitle, err := vo.NewTitle(b.title)
	if err != nil {
		return nil, err
	}

	// Validate and create markdown content value object
	validatedContent, err := vo.NewMarkdownContent(b.markdownContent)
	if err != nil {
		return nil, err
	}

	// Validate author
	if b.author == nil {
		return nil, errs.MissingAuthor
	}

	// Use provided slug or generate from title
	var slug *vo.Slug
	if b.slug != "" {
		slug = vo.NewSlug(b.slug)
	} else {
		slug = vo.NewSlugFromTitle(validatedTitle)
	}

	// Validate and create cover image URL value object
	validatedCoverURL, err := vo.NewCoverImageURL(b.coverImageURL)
	if err != nil {
		return nil, err
	}

	// Set timestamps if not provided
	now := time.Now()
	if b.createdAt.IsZero() {
		b.createdAt = now
	}
	if b.updatedAt.IsZero() {
		b.updatedAt = now
	}

	// Build and return the Post
	post := &Post{
		ID:              b.id,
		Title:           validatedTitle,
		Slug:            slug,
		Author:          b.author,
		CoverImageURL:   validatedCoverURL,
		MarkdownContent: validatedContent,
		CreatedAt:       b.createdAt,
		UpdatedAt:       b.updatedAt,
	}

	return post, nil
}
