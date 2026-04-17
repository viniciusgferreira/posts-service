package dto

import "github.com/viniciusgferreira/posts-service/internal/core/domain"

// CreatePostRequest represents the request payload for creating a new post
type CreatePostRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=200" example:"O Guia Completo para Arquitetura Hexagonal"`
	AuthorID        string `json:"author_id" binding:"required,min=1" example:"987654321"`
	CoverImageURL   string `json:"cover_image_url" binding:"omitempty,url" example:"https://cdn.seu-blog.com/imagens/post-arquitetura-hexagonal-capa.png"`
	MarkdownContent string `json:"markdown_content" binding:"required,min=1" example:"# Guia para Arquitetura Hexagonal..."`
}

// ToDomain converts CreatePostRequest to domain Post using the builder pattern
func (r *CreatePostRequest) ToDomain(author *domain.Author) (*domain.Post, error) {
	builder := domain.NewPostBuilder().
		WithTitle(r.Title).
		WithMarkdownContent(r.MarkdownContent).
		WithAuthor(author).
		WithCoverImageURL(r.CoverImageURL)

	return builder.Build()
}

// GetPostRequest represents the request parameters for getting a post by ID
type GetPostRequest struct {
	ID string `uri:"id" binding:"required,min=1" example:"123456789"`
}

// UpdatePostRequest represents the request payload for updating an existing post
type UpdatePostRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=200" example:"O Guia Completo para Arquitetura Hexagonal"`
	AuthorID        string `json:"author_id" binding:"required,min=1" example:"987654321"`
	CoverImageURL   string `json:"cover_image_url" binding:"omitempty,url" example:"https://cdn.seu-blog.com/imagens/post-arquitetura-hexagonal-capa.png"`
	MarkdownContent string `json:"markdown_content" binding:"required,min=1" example:"# Guia para Arquitetura Hexagonal..."`
}

// ToDomain converts UpdatePostRequest to domain Post using the builder pattern
func (r *UpdatePostRequest) ToDomain(author *domain.Author, postID string) (*domain.Post, error) {
	builder := domain.NewPostBuilder().
		WithID(postID).
		WithTitle(r.Title).
		WithMarkdownContent(r.MarkdownContent).
		WithAuthor(author).
		WithCoverImageURL(r.CoverImageURL)

	return builder.Build()
}
