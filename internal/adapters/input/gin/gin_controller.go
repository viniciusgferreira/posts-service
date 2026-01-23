package gin

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/viniciusgferreira/posts-service/internal/adapters/input/gin/dto"
	"github.com/viniciusgferreira/posts-service/internal/adapters/input/gin/hateoas"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
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
		c.handleError(ctx, errs.RequestBinding)
		return
	}

	// TODO: Implement create post logic using use cases
	// For now, return a mock response
	c.handleError(ctx, errors.New("Create post functionality not yet implemented"))
}

func (c *Controller) GetPost(ctx *gin.Context) {
	c.logger.Info("Get post requested")

	var req dto.GetPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.handleError(ctx, errs.RequestBinding)
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
	c.logger.Info("Update post requested")

	var uriReq dto.GetPostRequest
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		c.handleError(ctx, errs.RequestBinding)
		return
	}

	var req dto.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleError(ctx, errs.RequestBinding)
		return
	}

	c.logger.WithField("post_id", uriReq.ID).Info("Update post requested")

	// TODO: Implement update post logic using use cases
	// For now, return a mock response with HATEOAS links
	now := time.Now()
	slug := c.generateSlug(req.Title)
	baseURL := c.getBaseURL(ctx)

	hateoasBuilder := hateoas.NewBuilder(baseURL)
	links := hateoasBuilder.PostLinks(uriReq.ID, slug, req.AuthorID)

	response := dto.PostResponse{
		ID:              uriReq.ID,
		Title:           req.Title,
		Slug:            slug,
		AuthorID:        req.AuthorID,
		CoverImageURL:   req.CoverImageURL,
		MarkdownContent: req.MarkdownContent,
		CreatedAt:       time.Now().Add(-24 * time.Hour),
		UpdatedAt:       now,
		Links:           links,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) DeletePost(ctx *gin.Context) {
	c.logger.Info("Delete post requested")

	var req dto.GetPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.handleError(ctx, errs.RequestBinding)
		return
	}

	c.logger.WithField("post_id", req.ID).Info("Delete post requested")

	// TODO: Implement delete post logic using use cases
	ctx.Status(http.StatusNoContent)
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

// generateSlug creates a URL-friendly slug from a title string
func (c *Controller) generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Replace Portuguese characters
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

	// Keep only alphanumeric characters and hyphens
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	slug = result.String()

	// Remove multiple consecutive hyphens
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

// handleError handles HTTP errors by logging and returning a standardized error response
// It checks if the error is a custom app error and maps it appropriately
func (c *Controller) handleError(ctx *gin.Context, err error) {
	var appErr errs.AppErrorInterface
	if errors.As(err, &appErr) {
		// Custom app error detected
		statusCode := c.mapErrorTypeToStatusCode(appErr.GetType())

		c.logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"error_code":  appErr.GetCode(),
			"error_type":  appErr.GetType(),
			"message":     appErr.GetMessage(),
		}).Error("Request error")

		ctx.JSON(statusCode, dto.ErrorResponse{
			Error:     string(appErr.GetType()),
			Message:   appErr.GetMessage(),
			Code:      appErr.GetCode(),
			Timestamp: time.Now(),
		})
		return
	}

	// Fallback for non-custom errors
	c.logger.WithError(err).Error("Request error")
	ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
		Error:     string(errs.InternalType),
		Message:   err.Error(),
		Code:      errs.InternalServerError.GetCode(),
		Timestamp: time.Now(),
	})
}

// mapErrorTypeToStatusCode maps error types to HTTP status codes
func (c *Controller) mapErrorTypeToStatusCode(errorType errs.Type) int {
	switch errorType {
	case errs.ValidationType:
		return http.StatusBadRequest
	case errs.PermissionType:
		return http.StatusForbidden
	case errs.InternalType:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
