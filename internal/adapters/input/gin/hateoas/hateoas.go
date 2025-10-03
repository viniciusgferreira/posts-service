package hateoas

import (
	"fmt"
	"net/http"
	"strings"
)

// Link represents a HATEOAS link
type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Links represents a collection of HATEOAS links
type Links map[string]Link

// Builder helps build HATEOAS links
type Builder struct {
	baseURL string
}

// NewBuilder creates a new HATEOAS builder
func NewBuilder(baseURL string) *Builder {
	return &Builder{
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

// PostLinks creates links for a post resource
func (b *Builder) PostLinks(postID, slug, authorID string) Links {
	links := make(Links)

	// Self link
	links["self"] = Link{
		Rel:    "self",
		Href:   fmt.Sprintf("%s/posts/%s", b.baseURL, slug),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	// Edit link
	links["edit"] = Link{
		Rel:    "edit",
		Href:   fmt.Sprintf("%s/posts/%s", b.baseURL, slug),
		Method: http.MethodPut,
		Type:   "application/json",
	}

	// Delete link
	links["delete"] = Link{
		Rel:    "delete",
		Href:   fmt.Sprintf("%s/posts/%s", b.baseURL, slug),
		Method: http.MethodDelete,
		Type:   "application/json",
	}

	// Comments link
	links["comments"] = Link{
		Rel:    "comments",
		Href:   fmt.Sprintf("%s/posts/%s/comments", b.baseURL, slug),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	// Author link
	links["author"] = Link{
		Rel:    "author",
		Href:   fmt.Sprintf("%s/users/%s", b.baseURL, authorID),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	return links
}

// CollectionLinks creates links for a collection resource
func (b *Builder) CollectionLinks(resource string, page, limit, total int) Links {
	links := make(Links)

	// Self link
	links["self"] = Link{
		Rel:    "self",
		Href:   fmt.Sprintf("%s/%s?page=%d&limit=%d", b.baseURL, resource, page, limit),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	// First page
	if page > 1 {
		links["first"] = Link{
			Rel:    "first",
			Href:   fmt.Sprintf("%s/%s?page=1&limit=%d", b.baseURL, resource, limit),
			Method: http.MethodGet,
			Type:   "application/json",
		}
	}

	// Previous page
	if page > 1 {
		links["prev"] = Link{
			Rel:    "prev",
			Href:   fmt.Sprintf("%s/%s?page=%d&limit=%d", b.baseURL, resource, page-1, limit),
			Method: http.MethodGet,
			Type:   "application/json",
		}
	}

	// Next page
	totalPages := (total + limit - 1) / limit
	if page < totalPages {
		links["next"] = Link{
			Rel:    "next",
			Href:   fmt.Sprintf("%s/%s?page=%d&limit=%d", b.baseURL, resource, page+1, limit),
			Method: http.MethodGet,
			Type:   "application/json",
		}
	}

	// Last page
	if page < totalPages {
		links["last"] = Link{
			Rel:    "last",
			Href:   fmt.Sprintf("%s/%s?page=%d&limit=%d", b.baseURL, resource, totalPages, limit),
			Method: http.MethodGet,
			Type:   "application/json",
		}
	}

	// Create link
	links["create"] = Link{
		Rel:    "create",
		Href:   fmt.Sprintf("%s/%s", b.baseURL, resource),
		Method: http.MethodPost,
		Type:   "application/json",
	}

	return links
}

// UserLinks creates links for a user resource
func (b *Builder) UserLinks(userID string) Links {
	links := make(Links)

	// Self link
	links["self"] = Link{
		Rel:    "self",
		Href:   fmt.Sprintf("%s/users/%s", b.baseURL, userID),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	// Posts link
	links["posts"] = Link{
		Rel:    "posts",
		Href:   fmt.Sprintf("%s/posts?author_id=%s", b.baseURL, userID),
		Method: http.MethodGet,
		Type:   "application/json",
	}

	return links
}

// AddCustomLink adds a custom link to the collection
func (links Links) AddCustomLink(rel, href, method, contentType string) {
	links[rel] = Link{
		Rel:    rel,
		Href:   href,
		Method: method,
		Type:   contentType,
	}
}
