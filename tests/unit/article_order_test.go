package unit

import (
	"context"
	"testing"
	"time"

	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/repository"
)

func TestArticleAPIInterleavesSources(t *testing.T) {
	store := repository.NewMemoryStore()
	now := time.Now()
	articles := []models.Article{
		{ID: "gq-1", Source: "GQ", Title: "GQ 1", URL: "https://example.com/gq1", Summary: "s", ReadTime: "2 min read", PublishedAt: now.Add(5 * time.Minute), FetchedAt: now, Tags: []string{"Style"}},
		{ID: "gq-2", Source: "GQ", Title: "GQ 2", URL: "https://example.com/gq2", Summary: "s", ReadTime: "2 min read", PublishedAt: now.Add(4 * time.Minute), FetchedAt: now, Tags: []string{"Style"}},
		{ID: "esq-1", Source: "ESQUIRE", Title: "E 1", URL: "https://example.com/e1", Summary: "s", ReadTime: "2 min read", PublishedAt: now.Add(3 * time.Minute), FetchedAt: now, Tags: []string{"Style"}},
		{ID: "hs-1", Source: "HIGHSNOBIETY", Title: "H 1", URL: "https://example.com/h1", Summary: "s", ReadTime: "2 min read", PublishedAt: now.Add(2 * time.Minute), FetchedAt: now, Tags: []string{"Style"}},
	}
	if err := store.UpsertArticles(context.Background(), articles); err != nil {
		t.Fatal(err)
	}
	got, err := store.ListArticles(context.Background(), 4)
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < len(got); i++ {
		if got[i].Source == got[i-1].Source {
			t.Fatalf("adjacent sources were not interleaved: %+v", got)
		}
	}
}
