package dto

import (
	"time"

	"github.com/viniciusgferreira/posts-service/internal/adapters/input/gin/hateoas"
)

// CreatePostResponse represents the response payload for creating a new post
type PostResponse struct {
	ID              string        `json:"id" example:"123456789"`
	Title           string        `json:"title" example:"O Guia Completo para Arquitetura Hexagonal"`
	Slug            string        `json:"slug" example:"o-guia-completo-para-arquitetura-hexagonal"`
	AuthorID        string        `json:"author_id" example:"987654321"`
	CoverImageURL   string        `json:"cover_image_url" example:"https://cdn.seu-blog.com/imagens/post-arquitetura-hexagonal-capa.png"`
	MarkdownContent string        `json:"markdown_content" example:"# Guia para Arquitetura Hexagonal..."`
	CreatedAt       time.Time     `json:"created_at" example:"2025-09-03T10:00:00Z"`
	UpdatedAt       time.Time     `json:"updated_at" example:"2025-09-03T10:00:00Z"`
	Links           hateoas.Links `json:"_links"`
}

// GetPostsResponse represents the response for getting multiple posts
type GetPostsResponse struct {
	Posts      []PostResponse `json:"posts"`
	Pagination PaginationInfo `json:"pagination"`
	Links      hateoas.Links  `json:"_links"`
}

// PaginationInfo represents pagination metadata
type PaginationInfo struct {
	Page       int  `json:"page" example:"1"`
	Limit      int  `json:"limit" example:"10"`
	Total      int  `json:"total" example:"100"`
	TotalPages int  `json:"total_pages" example:"10"`
	HasNext    bool `json:"has_next" example:"true"`
	HasPrev    bool `json:"has_prev" example:"false"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"validation failed"`
	Message string `json:"message" example:"The title field is required"`
	Code    int    `json:"code" example:"400"`
}
