package domain

import (
	"errors"
	"strings"
	"time"
)

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAuthor(id, name, email string) (*Author, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("author name cannot be empty")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("author email cannot be empty")
	}

	now := time.Now()
	return &Author{
		ID:        id,
		Name:      strings.TrimSpace(name),
		Email:     strings.TrimSpace(email),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

type Post struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	Author          *Author   `json:"author"`
	CoverImageURL   string    `json:"cover_image_url"`
	MarkdownContent string    `json:"markdown_content"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewPost(id, title, markdownContent string, author *Author, coverImageURL string) (*Post, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("post title cannot be empty")
	}

	if strings.TrimSpace(markdownContent) == "" {
		return nil, errors.New("post content cannot be empty")
	}

	if author == nil {
		return nil, errors.New("post must have an author")
	}

	now := time.Now()
	post := &Post{
		ID:              id,
		Title:           strings.TrimSpace(title),
		Author:          author,
		CoverImageURL:   strings.TrimSpace(coverImageURL),
		MarkdownContent: strings.TrimSpace(markdownContent),
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	post.generateSlug()

	return post, nil
}

func (p *Post) generateSlug() {
	slug := strings.ToLower(p.Title)
	slug = strings.ReplaceAll(slug, " ", "-")

	slug = strings.ReplaceAll(slug, "ç", "c")
	slug = strings.ReplaceAll(slug, "ã", "a")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "à", "a")
	slug = strings.ReplaceAll(slug, "â", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "ê", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ô", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ü", "u")

	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	slug = result.String()

	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	slug = strings.Trim(slug, "-")

	p.Slug = slug
}

func (p *Post) UpdatePost(title, markdownContent, coverImageURL string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("post title cannot be empty")
	}

	if strings.TrimSpace(markdownContent) == "" {
		return errors.New("post content cannot be empty")
	}

	p.Title = strings.TrimSpace(title)
	p.MarkdownContent = strings.TrimSpace(markdownContent)
	p.CoverImageURL = strings.TrimSpace(coverImageURL)
	p.UpdatedAt = time.Now()

	p.generateSlug()

	return nil
}

func (p *Post) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("post title cannot be empty")
	}
	if strings.TrimSpace(p.MarkdownContent) == "" {
		return errors.New("post content cannot be empty")
	}
	if p.Author == nil {
		return errors.New("post must have an author")
	}
	if strings.TrimSpace(p.Slug) == "" {
		return errors.New("post slug cannot be empty")
	}
	return nil
}
