package hateoas

import (
	"fmt"
	"testing"
)

func ExampleBuilder_PostLinks() {
	// Create a HATEOAS builder
	builder := NewBuilder("https://api.example.com/v1")

	// Generate links for a post
	links := builder.PostLinks("123", "my-awesome-post", "456")

	// Print the links
	for rel, link := range links {
		fmt.Printf("%s: %s (%s)\n", rel, link.Href, link.Method)
	}

	// Output:
	// self: https://api.example.com/v1/posts/my-awesome-post (GET)
	// edit: https://api.example.com/v1/posts/my-awesome-post (PUT)
	// delete: https://api.example.com/v1/posts/my-awesome-post (DELETE)
	// comments: https://api.example.com/v1/posts/my-awesome-post/comments (GET)
	// author: https://api.example.com/v1/users/456 (GET)
}

func ExampleBuilder_CollectionLinks() {
	// Create a HATEOAS builder
	builder := NewBuilder("https://api.example.com/v1")

	// Generate collection links with pagination
	links := builder.CollectionLinks("posts", 2, 10, 25)

	// Print the links
	for rel, link := range links {
		fmt.Printf("%s: %s (%s)\n", rel, link.Href, link.Method)
	}

	// Output:
	// self: https://api.example.com/v1/posts?page=2&limit=10 (GET)
	// first: https://api.example.com/v1/posts?page=1&limit=10 (GET)
	// prev: https://api.example.com/v1/posts?page=1&limit=10 (GET)
	// next: https://api.example.com/v1/posts?page=3&limit=10 (GET)
	// last: https://api.example.com/v1/posts?page=3&limit=10 (GET)
	// create: https://api.example.com/v1/posts (POST)
}

func TestBuilder_PostLinks(t *testing.T) {
	builder := NewBuilder("https://api.example.com/v1")
	links := builder.PostLinks("123", "test-post", "456")

	// Check that all expected links are present
	expectedRels := []string{"self", "edit", "delete", "comments", "author"}
	for _, rel := range expectedRels {
		if _, exists := links[rel]; !exists {
			t.Errorf("Expected link relation '%s' not found", rel)
		}
	}

	// Check specific link
	if links["self"].Href != "https://api.example.com/v1/posts/test-post" {
		t.Errorf("Expected self link to be 'https://api.example.com/v1/posts/test-post', got '%s'", links["self"].Href)
	}
}
