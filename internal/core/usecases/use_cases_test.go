package usecases

import (
	"errors"
	"testing"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/errs"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
)

// --- Mock implementations for ports ---

type mockPostCreationPort struct {
	saveFunc func(post *domain.Post) error
}

func (m *mockPostCreationPort) Save(post *domain.Post) error {
	if m.saveFunc != nil {
		return m.saveFunc(post)
	}
	return nil
}

type mockAuthorReadingPort struct {
	findByIDFunc func(id string) (*domain.Author, error)
}

func (m *mockAuthorReadingPort) FindByID(id string) (*domain.Author, error) {
	if m.findByIDFunc != nil {
		return m.findByIDFunc(id)
	}
	return nil, nil
}

// --- Helpers ---

func mustEmail(email string) *vo.Email {
	e, err := vo.NewEmail(email)
	if err != nil {
		panic(err)
	}
	return e
}

func mustPost(t *testing.T) *domain.Post {
	t.Helper()
	author := &domain.Author{ID: "author-1", Name: "Test Author", Email: mustEmail("test@example.com")}
	post, err := domain.NewPostBuilder().
		WithTitle("Test Post Title").
		WithMarkdownContent("# Test Content with enough characters").
		WithAuthor(author).
		Build()
	if err != nil {
		t.Fatalf("failed to build test post: %v", err)
	}
	return post
}

// --- Tests ---

func TestPostUseCases_CreatePost_Success(t *testing.T) {
	saveCalled := false
	mockRepo := &mockPostCreationPort{
		saveFunc: func(post *domain.Post) error {
			saveCalled = true
			post.ID = "generated-id-123"
			return nil
		},
	}

	uc := NewPostUseCases(mockRepo, &mockAuthorReadingPort{})
	post := mustPost(t)

	result, err := uc.CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost() unexpected error: %v", err)
	}
	if !saveCalled {
		t.Error("expected Save to be called")
	}
	if result.ID != "generated-id-123" {
		t.Errorf("ID = %q, want %q", result.ID, "generated-id-123")
	}
}

func TestPostUseCases_CreatePost_DuplicateError(t *testing.T) {
	mockRepo := &mockPostCreationPort{
		saveFunc: func(post *domain.Post) error {
			return errs.DuplicateError
		},
	}

	uc := NewPostUseCases(mockRepo, &mockAuthorReadingPort{})
	post := mustPost(t)

	_, err := uc.CreatePost(post)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// The error should be wrapped but still matchable
	var appErr errs.AppErrorInterface
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppErrorInterface, got %T", err)
	}
}

func TestPostUseCases_GetAuthor_Success(t *testing.T) {
	expectedAuthor := &domain.Author{
		ID:    "author-1",
		Name:  "Test Author",
		Email: mustEmail("test@example.com"),
	}

	mockRepo := &mockAuthorReadingPort{
		findByIDFunc: func(id string) (*domain.Author, error) {
			if id != "author-1" {
				t.Errorf("FindByID called with %q, want %q", id, "author-1")
			}
			return expectedAuthor, nil
		},
	}

	uc := NewPostUseCases(&mockPostCreationPort{}, mockRepo)

	author, err := uc.GetAuthor("author-1")
	if err != nil {
		t.Fatalf("GetAuthor() unexpected error: %v", err)
	}
	if author.ID != expectedAuthor.ID {
		t.Errorf("Author.ID = %q, want %q", author.ID, expectedAuthor.ID)
	}
	if author.Name != expectedAuthor.Name {
		t.Errorf("Author.Name = %q, want %q", author.Name, expectedAuthor.Name)
	}
}

func TestPostUseCases_GetAuthor_NotFound(t *testing.T) {
	mockRepo := &mockAuthorReadingPort{
		findByIDFunc: func(id string) (*domain.Author, error) {
			return nil, errs.AuthorNotFound
		},
	}

	uc := NewPostUseCases(&mockPostCreationPort{}, mockRepo)

	_, err := uc.GetAuthor("nonexistent-id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var appErr errs.AppErrorInterface
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppErrorInterface, got %T: %v", err, err)
	}
	if appErr.GetCode() != 404 {
		t.Errorf("error code = %d, want 404", appErr.GetCode())
	}
}

func TestPostUseCases_GetAuthor_InfrastructureError(t *testing.T) {
	mockRepo := &mockAuthorReadingPort{
		findByIDFunc: func(id string) (*domain.Author, error) {
			return nil, errors.New("connection timeout")
		},
	}

	uc := NewPostUseCases(&mockPostCreationPort{}, mockRepo)

	_, err := uc.GetAuthor("author-1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Infrastructure errors should be wrapped, not be AppErrorInterface
	var appErr errs.AppErrorInterface
	if errors.As(err, &appErr) {
		t.Error("infrastructure error should not be AppErrorInterface")
	}
}
