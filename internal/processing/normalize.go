package processing

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"time"

	"fashion_dashboard/internal/models"
)

func HashCanonical(value string) string {
	sum := sha1.Sum([]byte(strings.TrimSpace(strings.ToLower(value))))
	return hex.EncodeToString(sum[:])
}

func NormalizeArticles(in []models.Article, now time.Time) []models.Article {
	out := make([]models.Article, 0, len(in))
	seen := map[string]bool{}
	for _, article := range in {
		if article.URL == "" || seen[article.URL] {
			continue
		}
		seen[article.URL] = true
		article.ContentHash = HashCanonical(article.URL)
		if article.ID == "" {
			article.ID = article.ContentHash
		}
		if article.FetchedAt.IsZero() {
			article.FetchedAt = now
		}
		if article.ReadTime == "" {
			article.ReadTime = "2 min read"
		}
		if len(article.Tags) == 0 {
			article.Tags = []string{"Style"}
		}
		out = append(out, article)
	}
	return out
}

func NormalizeLooks(in []models.Look, now time.Time) []models.Look {
	out := make([]models.Look, 0, len(in))
	seen := map[string]bool{}
	for _, look := range in {
		if look.SourceURL == "" || seen[look.SourceURL] {
			continue
		}
		seen[look.SourceURL] = true
		if look.ID == "" {
			look.ID = HashCanonical(look.SourceURL)
		}
		if look.FetchedAt.IsZero() {
			look.FetchedAt = now
		}
		if len(look.Tags) == 0 {
			look.Tags = []string{"Inspiration"}
		}
		out = append(out, look)
	}
	return out
}

func NormalizeItems(in []models.Item, now time.Time) []models.Item {
	out := make([]models.Item, 0, len(in))
	seen := map[string]bool{}
	for _, item := range in {
		if item.ProductURL == "" || seen[item.ProductURL] {
			continue
		}
		seen[item.ProductURL] = true
		if item.ID == "" {
			item.ID = HashCanonical(item.ProductURL)
		}
		if item.FetchedAt.IsZero() {
			item.FetchedAt = now
		}
		if len(item.Tags) == 0 {
			item.Tags = []string{item.Category}
		}
		out = append(out, item)
	}
	return out
}
