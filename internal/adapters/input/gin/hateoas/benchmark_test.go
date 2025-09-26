package hateoas

import (
	"encoding/json"
	"net/http"
	"testing"
)

func BenchmarkBuilder_PostLinks(b *testing.B) {
	builder := NewBuilder("https://api.example.com/v1")
	postID := "123"
	slug := "my-awesome-post"
	authorID := "456"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = builder.PostLinks(postID, slug, authorID)
	}
}

func BenchmarkBuilder_CollectionLinks(b *testing.B) {
	builder := NewBuilder("https://api.example.com/v1")
	resource := "posts"
	page := 2
	limit := 10
	total := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = builder.CollectionLinks(resource, page, limit, total)
	}
}

func BenchmarkBuilder_UserLinks(b *testing.B) {
	builder := NewBuilder("https://api.example.com/v1")
	userID := "123"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = builder.UserLinks(userID)
	}
}

func BenchmarkLinks_AddCustomLink(b *testing.B) {
	links := make(Links)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		links.AddCustomLink("custom", "/custom/path", http.MethodPost, "application/json")
	}
}

func BenchmarkLinks_HasLink(b *testing.B) {
	links := make(Links)
	links["self"] = Link{Rel: "self", Href: "/posts/123"}
	links["edit"] = Link{Rel: "edit", Href: "/posts/123"}
	links["delete"] = Link{Rel: "delete", Href: "/posts/123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = links.HasLink("self")
	}
}

func BenchmarkLink_JSONMarshal(b *testing.B) {
	link := Link{
		Rel:    "self",
		Href:   "https://api.example.com/posts/123",
		Method: http.MethodGet,
		Type:   "application/json",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(link)
	}
}

func BenchmarkLink_JSONUnmarshal(b *testing.B) {
	jsonData := []byte(`{"rel":"self","href":"https://api.example.com/posts/123","method":"GET","type":"application/json"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var link Link
		_ = json.Unmarshal(jsonData, &link)
	}
}
