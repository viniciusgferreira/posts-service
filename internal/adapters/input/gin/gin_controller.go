package gin

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/viniciusgferreira/posts-service/internal/adapters/input/gin/dto"
	"github.com/viniciusgferreira/posts-service/internal/adapters/input/gin/hateoas"
)

type Controller struct {
	logger *logrus.Logger
}

func NewController(logger *logrus.Logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

func (c *Controller) SetupRoutes(r *gin.Engine) {
	r.GET("/health", c.HealthCheck)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/posts", c.GetPosts)
		v1.POST("/posts", c.CreatePost)
		v1.GET("/posts/:id", c.GetPost)
		v1.PUT("/posts/:id", c.UpdatePost)
		v1.DELETE("/posts/:id", c.DeletePost)
	}
}

func (c *Controller) HealthCheck(ctx *gin.Context) {
	c.logger.Info("Health check requested")
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "posts-service",
	})
}

func (c *Controller) GetPosts(ctx *gin.Context) {
	c.logger.Info("Get posts requested")

	page := 1
	limit := 10
	// TODO: Use these query parameters in the actual implementation
	_ = ctx.Query("author_id")
	_ = ctx.Query("status")
	_ = ctx.Query("search")

	// TODO: Implement get posts logic using use cases
	// For now, return a mock response with HATEOAS links
	baseURL := c.getBaseURL(ctx)
	hateoasBuilder := hateoas.NewBuilder(baseURL)

	// Mock data
	posts := []dto.PostResponse{
		{
			ID:              "123456789",
			Title:           "O Guia Completo para Arquitetura Hexagonal",
			Slug:            "o-guia-completo-para-arquitetura-hexagonal",
			AuthorID:        "987654321",
			CoverImageURL:   "https://cdn.seu-blog.com/imagens/post-arquitetura-hexagonal-capa.png",
			MarkdownContent: "# Guia para Arquitetura Hexagonal...",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			Links:           hateoasBuilder.PostLinks("123456789", "o-guia-completo-para-arquitetura-hexagonal", "987654321"),
		},
	}

	collectionLinks := hateoasBuilder.CollectionLinks("posts", page, limit, 1)

	response := dto.GetPostsResponse{
		Posts: posts,
		Pagination: dto.PaginationInfo{
			Page:       page,
			Limit:      limit,
			Total:      1,
			TotalPages: 1,
			HasNext:    false,
			HasPrev:    false,
		},
		Links: collectionLinks,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) CreatePost(ctx *gin.Context) {
	c.logger.Info("Create post requested")

	var req dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.WithError(err).Error("Failed to bind request body")
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation failed",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// TODO: Implement create post logic using use cases
	// For now, return a mock response with HATEOAS links
	now := time.Now()
	postID := "123456789"
	slug := c.generateSlug(req.Title)
	baseURL := c.getBaseURL(ctx)

	hateoasBuilder := hateoas.NewBuilder(baseURL)
	links := hateoasBuilder.PostLinks(postID, slug, req.AuthorID)

	response := dto.PostResponse{
		ID:              postID,
		Title:           req.Title,
		Slug:            slug,
		AuthorID:        req.AuthorID,
		CoverImageURL:   req.CoverImageURL,
		MarkdownContent: req.MarkdownContent,
		CreatedAt:       now,
		UpdatedAt:       now,
		Links:           links,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) GetPost(ctx *gin.Context) {
	c.logger.Info("Get post requested")

	var req dto.GetPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.WithError(err).Error("Failed to bind URI parameters")
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation failed",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	c.logger.WithField("post_id", req.ID).Info("Get post requested")

	// TODO: Implement get post logic using use cases
	// Placeholder response with proper structure and HATEOAS links
	baseURL := c.getBaseURL(ctx)
	hateoasBuilder := hateoas.NewBuilder(baseURL)

	// Generate HATEOAS links for the post
	links := hateoasBuilder.PostLinks(req.ID, "placeholder-slug", "placeholder-author-id")

	response := dto.PostResponse{
		ID:              req.ID,
		Title:           "Post Title (Placeholder)",
		Slug:            "placeholder-slug",
		AuthorID:        "placeholder-author-id",
		CoverImageURL:   "",
		MarkdownContent: "# Placeholder Post\n\nThis is a placeholder response. The actual post data will be fetched from the database when the use case is implemented.",
		CreatedAt:       time.Now().Add(-24 * time.Hour), // Created 1 day ago
		UpdatedAt:       time.Now(),
		Links:           links,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	c.logger.WithField("post_id", id).Info("Update post requested")
	// TODO: Implement update post logic
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update post endpoint - to be implemented",
		"id":      id,
	})
}

func (c *Controller) DeletePost(ctx *gin.Context) {
	c.logger.Info("Delete post requested")

	var req dto.GetPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.WithError(err).Error("Failed to bind URI parameters")
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation failed",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	c.logger.WithField("post_id", req.ID).Info("Delete post requested")

	// TODO: Implement delete post logic using use cases
	ctx.Status(http.StatusNoContent)
}

func (c *Controller) generateSlug(title string) string {
	slug := strings.ToLower(title)
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
	return slug
}

func (c *Controller) getBaseURL(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	host := ctx.Request.Host
	if host == "" {
		host = "localhost:8080"
	}

	return scheme + "://" + host + "/api/v1"
}
