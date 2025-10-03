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
	// Health check endpoint
	r.GET("/health", c.HealthCheck)

	// API v1 routes
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

	// Parse query parameters
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
	posts := []dto.CreatePostResponse{
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

	// Generate collection links
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

	// Create HATEOAS builder and generate links
	hateoasBuilder := hateoas.NewBuilder(baseURL)
	links := hateoasBuilder.PostLinks(postID, slug, req.AuthorID)

	response := dto.CreatePostResponse{
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
	id := ctx.Param("id")
	c.logger.WithField("post_id", id).Info("Get post requested")
	// TODO: Implement get post logic
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get post endpoint - to be implemented",
		"id":      id,
	})
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
	id := ctx.Param("id")
	c.logger.WithField("post_id", id).Info("Delete post requested")
	// TODO: Implement delete post logic
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete post endpoint - to be implemented",
		"id":      id,
	})
}

// generateSlug creates a URL-friendly slug from a title
func (c *Controller) generateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters except hyphens
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

	// Remove any remaining special characters except hyphens and alphanumeric
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	// Remove multiple consecutive hyphens
	slug = result.String()
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

// getBaseURL extracts the base URL from the request context
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
