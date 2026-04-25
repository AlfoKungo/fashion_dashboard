package unit

import (
	"time"

	"fashion_dashboard/internal/models"
)

func article(id string) models.Article {
	return models.Article{ID: id, Source: "GQ", Title: "Title", URL: "https://example.com/" + id, Summary: "Summary", ReadTime: "2 min read", Tags: []string{"Tag"}, FetchedAt: time.Now()}
}

func look(id string) models.Look {
	return models.Look{ID: id, Source: "Editorial", Title: "Look", SourceURL: "https://example.com/" + id, Tags: []string{"Tag"}, FetchedAt: time.Now()}
}

func item(id, category string) models.Item {
	return models.Item{ID: id, Source: "Retail", Brand: "Brand", Name: "Item", Category: category, Price: "$100", ProductURL: "https://example.com/" + id, Tags: []string{category}, FetchedAt: time.Now()}
}
