package unit

import (
	"context"
	"testing"

	"fashion_dashboard/internal/repository"
)

func TestMemoryStoreProvidesSeedData(t *testing.T) {
	store := repository.NewMemoryStore()
	articles, err := store.ListArticles(context.Background(), 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(articles) != 4 {
		t.Fatalf("articles = %d", len(articles))
	}
}
