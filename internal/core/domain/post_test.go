package domain

import (
	"testing"

	"github.com/viniciusgferreira/posts-service/internal/core/domain/vo"
)

func TestPostBuilder_SlugFromTitle(t *testing.T) {
	author := &Author{ID: "author-1", Name: "Test Author", Email: mustEmail("test@example.com")}

	post, err := NewPostBuilder().
		WithTitle("My Awesome Post").
		WithMarkdownContent("# Content with enough characters").
		WithAuthor(author).
		Build()

	if err != nil {
		t.Fatalf("Build() unexpected error: %v", err)
	}

	expectedSlug := "my-awesome-post"
	if post.Slug.String() != expectedSlug {
		t.Errorf("Slug = %q, want %q (derived from title)", post.Slug.String(), expectedSlug)
	}
}

func TestPostBuilder_ExplicitSlug(t *testing.T) {
	author := &Author{ID: "author-1", Name: "Test Author", Email: mustEmail("test@example.com")}

	post, err := NewPostBuilder().
		WithTitle("My Awesome Post").
		WithSlug("custom-slug").
		WithMarkdownContent("# Content with enough characters").
		WithAuthor(author).
		Build()

	if err != nil {
		t.Fatalf("Build() unexpected error: %v", err)
	}

	expectedSlug := "custom-slug"
	if post.Slug.String() != expectedSlug {
		t.Errorf("Slug = %q, want %q (explicit slug)", post.Slug.String(), expectedSlug)
	}
}

func TestPostBuilder_EmptySlugFallsBackToTitle(t *testing.T) {
	author := &Author{ID: "author-1", Name: "Test Author", Email: mustEmail("test@example.com")}

	post, err := NewPostBuilder().
		WithTitle("Portuguese Title Ação").
		WithSlug("").
		WithMarkdownContent("# Content with enough characters").
		WithAuthor(author).
		Build()

	if err != nil {
		t.Fatalf("Build() unexpected error: %v", err)
	}

	expectedSlug := "portuguese-title-acao"
	if post.Slug.String() != expectedSlug {
		t.Errorf("Slug = %q, want %q (derived from title when slug is empty)", post.Slug.String(), expectedSlug)
	}
}

func mustEmail(email string) *vo.Email {
	e, err := vo.NewEmail(email)
	if err != nil {
		panic(err)
	}
	return e
}
