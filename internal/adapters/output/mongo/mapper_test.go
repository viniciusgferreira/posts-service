package mongo

import (
	"testing"
	"time"

	"github.com/viniciusgferreira/posts-service/internal/core/domain"
	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mustEmail(email string) *vo.Email {
	e, err := vo.NewEmail(email)
	if err != nil {
		panic(err)
	}
	return e
}

func TestToDocument_NewPost(t *testing.T) {
	author := &domain.Author{ID: "author-123", Name: "Test Author", Email: mustEmail("test@example.com")}
	post, err := domain.NewPostBuilder().
		WithTitle("My Test Post").
		WithMarkdownContent("# Content with enough characters here").
		WithAuthor(author).
		WithCoverImageURL("https://example.com/image.png").
		Build()
	if err != nil {
		t.Fatalf("failed to build post: %v", err)
	}

	doc := toDocument(post)

	if doc.Title != "My Test Post" {
		t.Errorf("Title = %q, want %q", doc.Title, "My Test Post")
	}
	if doc.Slug != "my-test-post" {
		t.Errorf("Slug = %q, want %q", doc.Slug, "my-test-post")
	}
	if doc.AuthorID != "author-123" {
		t.Errorf("AuthorID = %q, want %q", doc.AuthorID, "author-123")
	}
	if doc.AuthorName != "Test Author" {
		t.Errorf("AuthorName = %q, want %q", doc.AuthorName, "Test Author")
	}
	if doc.AuthorEmail != "test@example.com" {
		t.Errorf("AuthorEmail = %q, want %q", doc.AuthorEmail, "test@example.com")
	}
	if doc.CoverImageURL != "https://example.com/image.png" {
		t.Errorf("CoverImageURL = %q, want %q", doc.CoverImageURL, "https://example.com/image.png")
	}
	if doc.MarkdownContent != "# Content with enough characters here" {
		t.Errorf("MarkdownContent = %q, want %q", doc.MarkdownContent, "# Content with enough characters here")
	}
	// New post without ID should have zero ObjectID
	if doc.ID != primitive.NilObjectID {
		t.Errorf("ID should be NilObjectID for new post, got %v", doc.ID)
	}
}

func TestToDocument_ExistingPost(t *testing.T) {
	author := &domain.Author{ID: "author-123", Name: "Test Author", Email: mustEmail("test@example.com")}
	validID := primitive.NewObjectID().Hex()

	post, err := domain.NewPostBuilder().
		WithID(validID).
		WithTitle("Existing Post").
		WithMarkdownContent("# Content with enough characters here").
		WithAuthor(author).
		Build()
	if err != nil {
		t.Fatalf("failed to build post: %v", err)
	}

	doc := toDocument(post)

	if doc.ID.Hex() != validID {
		t.Errorf("ID = %q, want %q", doc.ID.Hex(), validID)
	}
}

func TestToDomain_ValidDocument(t *testing.T) {
	objID := primitive.NewObjectID()
	now := time.Now().Truncate(time.Millisecond)

	doc := postDocument{
		ID:              objID,
		Title:           "Test Post",
		Slug:            "test-post",
		AuthorID:        "author-456",
		AuthorName:      "Author Name",
		AuthorEmail:     "author@example.com",
		CoverImageURL:   "https://example.com/cover.jpg",
		MarkdownContent: "# Hello World with enough content",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	post, err := toDomain(doc)
	if err != nil {
		t.Fatalf("toDomain() unexpected error: %v", err)
	}

	if post.ID != objID.Hex() {
		t.Errorf("ID = %q, want %q", post.ID, objID.Hex())
	}
	if post.Title.String() != "Test Post" {
		t.Errorf("Title = %q, want %q", post.Title.String(), "Test Post")
	}
	if post.Slug.String() != "test-post" {
		t.Errorf("Slug = %q, want %q", post.Slug.String(), "test-post")
	}
	if post.Author.ID != "author-456" {
		t.Errorf("Author.ID = %q, want %q", post.Author.ID, "author-456")
	}
	if post.Author.Name != "Author Name" {
		t.Errorf("Author.Name = %q, want %q", post.Author.Name, "Author Name")
	}
}

func TestToDomain_InvalidEmail(t *testing.T) {
	doc := postDocument{
		ID:              primitive.NewObjectID(),
		Title:           "Test",
		Slug:            "test",
		AuthorID:        "a1",
		AuthorName:      "Name",
		AuthorEmail:     "invalid-email",
		MarkdownContent: "# content enough chars",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	_, err := toDomain(doc)
	if err == nil {
		t.Fatal("expected error for invalid email, got nil")
	}
}

func TestToAuthorDomain_Valid(t *testing.T) {
	objID := primitive.NewObjectID()
	doc := authorDocument{
		ID:    objID,
		Name:  "Author",
		Email: "author@example.com",
	}

	author, err := toAuthorDomain(doc)
	if err != nil {
		t.Fatalf("toAuthorDomain() unexpected error: %v", err)
	}

	if author.ID != objID.Hex() {
		t.Errorf("ID = %q, want %q", author.ID, objID.Hex())
	}
	if author.Name != "Author" {
		t.Errorf("Name = %q, want %q", author.Name, "Author")
	}
	if author.Email.String() != "author@example.com" {
		t.Errorf("Email = %q, want %q", author.Email.String(), "author@example.com")
	}
}

func TestToAuthorDomain_InvalidEmail(t *testing.T) {
	doc := authorDocument{
		ID:    primitive.NewObjectID(),
		Name:  "Author",
		Email: "not-an-email",
	}

	_, err := toAuthorDomain(doc)
	if err == nil {
		t.Fatal("expected error for invalid email, got nil")
	}
}
