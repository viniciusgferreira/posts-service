package dto

// CreatePostRequest represents the request payload for creating a new post
type CreatePostRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=200" example:"O Guia Completo para Arquitetura Hexagonal"`
	AuthorID        string `json:"author_id" binding:"required,min=1" example:"987654321"`
	CoverImageURL   string `json:"cover_image_url" binding:"omitempty,url" example:"https://cdn.seu-blog.com/imagens/post-arquitetura-hexagonal-capa.png"`
	MarkdownContent string `json:"markdown_content" binding:"required,min=1" example:"# Guia para Arquitetura Hexagonal..."`
}
