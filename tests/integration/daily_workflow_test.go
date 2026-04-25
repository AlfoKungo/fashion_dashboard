package integration

import (
	"context"
	"testing"
	"time"

	articlefetcher "fashion_dashboard/internal/fetchers/articles"
	itemfetcher "fashion_dashboard/internal/fetchers/items"
	lookfetcher "fashion_dashboard/internal/fetchers/looks"
	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/scheduler"
)

func TestDailyWorkflow(t *testing.T) {
	store := repository.NewMemoryStore()
	workflow := scheduler.NewWorkflowWithFetchers(store, fakeArticles{}, lookfetcher.SampleFetcher{}, itemfetcher.SampleFetcher{})
	if err := workflow.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	looks, err := store.ListDailyLooks(context.Background(), "", 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(looks) == 0 {
		t.Fatalf("expected daily looks")
	}
}

type fakeArticles struct{}

func (fakeArticles) Fetch(context.Context) ([]models.Article, error) {
	now := time.Now()
	return []models.Article{
		{ID: "article-a", Source: "TEST", Title: "A", URL: "https://example.com/a", Summary: "A summary", ReadTime: "2 min read", Tags: []string{"Style"}, FetchedAt: now},
	}, nil
}

var _ articlefetcher.Fetcher = fakeArticles{}
