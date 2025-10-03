package hateoas

import (
	"encoding/json"
	"net/http"
	"sort"
	"testing"
)

func TestLink_JSONSerialization(t *testing.T) {
	tests := []struct {
		name     string
		link     Link
		expected string
	}{
		{
			name: "Complete link with all fields",
			link: Link{
				Rel:    "self",
				Href:   "https://api.example.com/posts/123",
				Method: http.MethodGet,
				Type:   "application/json",
			},
			expected: `{"rel":"self","href":"https://api.example.com/posts/123","method":"GET","type":"application/json"}`,
		},
		{
			name: "Link with minimal fields",
			link: Link{
				Rel:  "self",
				Href: "https://api.example.com/posts/123",
			},
			expected: `{"rel":"self","href":"https://api.example.com/posts/123"}`,
		},
		{
			name:     "Link with empty fields",
			link:     Link{},
			expected: `{"rel":"","href":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test JSON marshaling
			jsonBytes, err := json.Marshal(tt.link)
			if err != nil {
				t.Fatalf("Failed to marshal link to JSON: %v", err)
			}

			if string(jsonBytes) != tt.expected {
				t.Errorf("Expected JSON: %s, got: %s", tt.expected, string(jsonBytes))
			}

			// Test JSON unmarshaling
			var unmarshaledLink Link
			err = json.Unmarshal(jsonBytes, &unmarshaledLink)
			if err != nil {
				t.Fatalf("Failed to unmarshal JSON to link: %v", err)
			}

			if unmarshaledLink != tt.link {
				t.Errorf("Expected link: %+v, got: %+v", tt.link, unmarshaledLink)
			}
		})
	}
}

func TestLink_Validation(t *testing.T) {
	tests := []struct {
		name        string
		link        Link
		expectValid bool
	}{
		{
			name: "Valid complete link",
			link: Link{
				Rel:    "self",
				Href:   "https://api.example.com/posts/123",
				Method: http.MethodGet,
				Type:   "application/json",
			},
			expectValid: true,
		},
		{
			name: "Valid minimal link",
			link: Link{
				Rel:  "self",
				Href: "https://api.example.com/posts/123",
			},
			expectValid: true,
		},
		{
			name: "Invalid link - empty rel",
			link: Link{
				Rel:  "",
				Href: "https://api.example.com/posts/123",
			},
			expectValid: false,
		},
		{
			name: "Invalid link - empty href",
			link: Link{
				Rel:  "self",
				Href: "",
			},
			expectValid: false,
		},
		{
			name: "Invalid link - empty rel and href",
			link: Link{
				Rel:  "",
				Href: "",
			},
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := isValidLink(tt.link)
			if isValid != tt.expectValid {
				t.Errorf("Expected valid: %v, got: %v", tt.expectValid, isValid)
			}
		})
	}
}

func TestLinks_Operations(t *testing.T) {
	t.Run("Add and retrieve links", func(t *testing.T) {
		links := make(Links)

		// Add a link
		link := Link{
			Rel:    "self",
			Href:   "https://api.example.com/posts/123",
			Method: http.MethodGet,
			Type:   "application/json",
		}
		links["self"] = link

		// Retrieve the link
		retrievedLink, exists := links["self"]
		if !exists {
			t.Fatal("Expected link to exist")
		}

		if retrievedLink != link {
			t.Errorf("Expected link: %+v, got: %+v", link, retrievedLink)
		}
	})

	t.Run("Add custom link", func(t *testing.T) {
		links := make(Links)
		links.AddCustomLink("publish", "/posts/123/publish", http.MethodPost, "application/json")

		customLink, exists := links["publish"]
		if !exists {
			t.Fatal("Expected custom link to exist")
		}

		expectedLink := Link{
			Rel:    "publish",
			Href:   "/posts/123/publish",
			Method: http.MethodPost,
			Type:   "application/json",
		}

		if customLink != expectedLink {
			t.Errorf("Expected custom link: %+v, got: %+v", expectedLink, customLink)
		}
	})

	t.Run("Check link existence", func(t *testing.T) {
		links := make(Links)
		links["self"] = Link{Rel: "self", Href: "/posts/123"}

		if !links.HasLink("self") {
			t.Error("Expected link 'self' to exist")
		}

		if links.HasLink("nonexistent") {
			t.Error("Expected link 'nonexistent' to not exist")
		}
	})

	t.Run("Get all link relations", func(t *testing.T) {
		links := make(Links)
		links["self"] = Link{Rel: "self", Href: "/posts/123"}
		links["edit"] = Link{Rel: "edit", Href: "/posts/123"}
		links["delete"] = Link{Rel: "delete", Href: "/posts/123"}

		relations := links.GetRelations()
		expectedRelations := []string{"delete", "edit", "self"} // Should be sorted

		if len(relations) != len(expectedRelations) {
			t.Errorf("Expected %d relations, got %d", len(expectedRelations), len(relations))
		}

		for i, rel := range relations {
			if rel != expectedRelations[i] {
				t.Errorf("Expected relation at index %d: %s, got: %s", i, expectedRelations[i], rel)
			}
		}
	})
}

func TestBuilder_NewBuilder(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		expected string
	}{
		{
			name:     "URL without trailing slash",
			baseURL:  "https://api.example.com/v1",
			expected: "https://api.example.com/v1",
		},
		{
			name:     "URL with trailing slash",
			baseURL:  "https://api.example.com/v1/",
			expected: "https://api.example.com/v1",
		},
		{
			name:     "URL with multiple trailing slashes",
			baseURL:  "https://api.example.com/v1///",
			expected: "https://api.example.com/v1",
		},
		{
			name:     "Empty URL",
			baseURL:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewBuilder(tt.baseURL)
			if builder.baseURL != tt.expected {
				t.Errorf("Expected baseURL: %s, got: %s", tt.expected, builder.baseURL)
			}
		})
	}
}

func TestBuilder_PostLinks(t *testing.T) {
	builder := NewBuilder("https://api.example.com/v1")

	tests := []struct {
		name     string
		postID   string
		slug     string
		authorID string
	}{
		{
			name:     "Valid post links",
			postID:   "123",
			slug:     "my-awesome-post",
			authorID: "456",
		},
		{
			name:     "Post with special characters in slug",
			postID:   "123",
			slug:     "post-with-special-chars-!@#",
			authorID: "456",
		},
		{
			name:     "Empty IDs",
			postID:   "",
			slug:     "",
			authorID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links := builder.PostLinks(tt.postID, tt.slug, tt.authorID)

			// Check that all expected links are present
			expectedRels := []string{"self", "edit", "delete", "comments", "author"}
			for _, rel := range expectedRels {
				if !links.HasLink(rel) {
					t.Errorf("Expected link relation '%s' not found", rel)
				}
			}

			// Check specific link structures
			selfLink := links["self"]
			if selfLink.Rel != "self" {
				t.Errorf("Expected self link rel to be 'self', got: %s", selfLink.Rel)
			}
			if selfLink.Method != http.MethodGet {
				t.Errorf("Expected self link method to be GET, got: %s", selfLink.Method)
			}
			if selfLink.Type != "application/json" {
				t.Errorf("Expected self link type to be 'application/json', got: %s", selfLink.Type)
			}

			// Check href format
			expectedSelfHref := "https://api.example.com/v1/posts/" + tt.slug
			if selfLink.Href != expectedSelfHref {
				t.Errorf("Expected self link href: %s, got: %s", expectedSelfHref, selfLink.Href)
			}

			// Check author link
			authorLink := links["author"]
			expectedAuthorHref := "https://api.example.com/v1/users/" + tt.authorID
			if authorLink.Href != expectedAuthorHref {
				t.Errorf("Expected author link href: %s, got: %s", expectedAuthorHref, authorLink.Href)
			}
		})
	}
}

func TestBuilder_CollectionLinks(t *testing.T) {
	builder := NewBuilder("https://api.example.com/v1")

	tests := []struct {
		name       string
		resource   string
		page       int
		limit      int
		total      int
		expectNext bool
		expectPrev bool
	}{
		{
			name:       "First page with next",
			resource:   "posts",
			page:       1,
			limit:      10,
			total:      25,
			expectNext: true,
			expectPrev: false,
		},
		{
			name:       "Middle page with both",
			resource:   "posts",
			page:       2,
			limit:      10,
			total:      25,
			expectNext: true,
			expectPrev: true,
		},
		{
			name:       "Last page with prev",
			resource:   "posts",
			page:       3,
			limit:      10,
			total:      25,
			expectNext: false,
			expectPrev: true,
		},
		{
			name:       "Single page",
			resource:   "posts",
			page:       1,
			limit:      10,
			total:      5,
			expectNext: false,
			expectPrev: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links := builder.CollectionLinks(tt.resource, tt.page, tt.limit, tt.total)

			// Check that self link is always present
			if !links.HasLink("self") {
				t.Error("Expected 'self' link to be present")
			}

			// Check self link href format
			selfLink := links["self"]
			expectedSelfHref := "https://api.example.com/v1/posts?page=1&limit=10"
			if tt.page == 1 && tt.limit == 10 {
				if selfLink.Href != expectedSelfHref {
					t.Errorf("Expected self link href: %s, got: %s", expectedSelfHref, selfLink.Href)
				}
			}

			// Check next link
			if tt.expectNext {
				if !links.HasLink("next") {
					t.Error("Expected 'next' link to be present")
				}
			} else {
				if links.HasLink("next") {
					t.Error("Expected 'next' link to not be present")
				}
			}

			// Check prev link
			if tt.expectPrev {
				if !links.HasLink("prev") {
					t.Error("Expected 'prev' link to be present")
				}
			} else {
				if links.HasLink("prev") {
					t.Error("Expected 'prev' link to not be present")
				}
			}

			// Check create link is always present
			if !links.HasLink("create") {
				t.Error("Expected 'create' link to be present")
			}

			createLink := links["create"]
			if createLink.Method != http.MethodPost {
				t.Errorf("Expected create link method to be POST, got: %s", createLink.Method)
			}
		})
	}
}

func TestBuilder_UserLinks(t *testing.T) {
	builder := NewBuilder("https://api.example.com/v1")

	tests := []struct {
		name   string
		userID string
	}{
		{
			name:   "Valid user ID",
			userID: "123",
		},
		{
			name:   "Empty user ID",
			userID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links := builder.UserLinks(tt.userID)

			// Check that self link is present
			if !links.HasLink("self") {
				t.Error("Expected 'self' link to be present")
			}

			selfLink := links["self"]
			expectedSelfHref := "https://api.example.com/v1/users/" + tt.userID
			if selfLink.Href != expectedSelfHref {
				t.Errorf("Expected self link href: %s, got: %s", expectedSelfHref, selfLink.Href)
			}

			// Check posts link is present
			if !links.HasLink("posts") {
				t.Error("Expected 'posts' link to be present")
			}

			postsLink := links["posts"]
			expectedPostsHref := "https://api.example.com/v1/posts?author_id=" + tt.userID
			if postsLink.Href != expectedPostsHref {
				t.Errorf("Expected posts link href: %s, got: %s", expectedPostsHref, postsLink.Href)
			}
		})
	}
}

// Helper functions for testing

func isValidLink(link Link) bool {
	return link.Rel != "" && link.Href != ""
}

func (links Links) HasLink(rel string) bool {
	_, exists := links[rel]
	return exists
}

func (links Links) GetRelations() []string {
	relations := make([]string, 0, len(links))
	for rel := range links {
		relations = append(relations, rel)
	}
	// Sort for consistent testing
	sort.Strings(relations)
	return relations
}
