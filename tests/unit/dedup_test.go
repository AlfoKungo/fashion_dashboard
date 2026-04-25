package unit

import (
	"testing"
	"time"

	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/processing"
)

func TestNormalizeArticlesDeduplicatesURLs(t *testing.T) {
	items := []models.Article{article("a"), article("a")}
	got := processing.NormalizeArticles(items, time.Now())
	if len(got) != 1 {
		t.Fatalf("expected 1 article, got %d", len(got))
	}
}
