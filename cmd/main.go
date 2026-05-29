package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginadapter "github.com/viniciusgferreira/posts-service/internal/adapters/input/gin"
	mongoadapter "github.com/viniciusgferreira/posts-service/internal/adapters/output/mongo"
	"github.com/viniciusgferreira/posts-service/internal/config"
	"github.com/viniciusgferreira/posts-service/internal/core/usecases"
)

const shutdownTimeout = 30 * time.Second

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.GinMode != "" {
		gin.SetMode(cfg.GinMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to MongoDB
	mongoClient, db, err := mongoadapter.NewConnection(cfg.MongoDB.URI, cfg.MongoDB.Database, logger)
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to MongoDB")
	}
	defer mongoadapter.Disconnect(mongoClient, logger)

	// Initialize repositories (output adapters)
	postRepository := mongoadapter.NewPostRepository(db)
	authorRepository := mongoadapter.NewAuthorRepository(db)

	// Initialize use cases
	postUseCases := usecases.NewPostUseCases(postRepository, authorRepository)

	// Create Gin router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Initialize controller with use cases
	controller := ginadapter.NewController(logger, postUseCases)
	controller.SetupRoutes(router)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.WithField("port", cfg.Port).Info("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	logger.Info("Server exited")
}

// corsMiddleware adds CORS headers to all responses
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
