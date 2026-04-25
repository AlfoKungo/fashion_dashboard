package repository

import (
	"fmt"
	"time"

	"fashion_dashboard/internal/models"
)

func seedArticles(now time.Time) []models.Article {
	sources := []string{"GQ", "HIGHSNOBIETY", "MR PORTER", "ESQUIRE"}
	titles := []string{"The Return of Soft Tailoring", "Footwear That Works All Day", "How to Layer Linen", "The New Neutral Palette"}
	tags := [][]string{{"Tailoring"}, {"Footwear"}, {"Summer"}, {"Trend Analysis"}}
	out := make([]models.Article, 0, 4)
	for i := range sources {
		id := fmt.Sprintf("article-%d", i+1)
		out = append(out, models.Article{
			ID:          id,
			Source:      sources[i],
			Title:       titles[i],
			URL:         "https://example.com/articles/" + id,
			ImageURL:    "https://picsum.photos/seed/" + id + "/900/600",
			Summary:     "A concise briefing on silhouettes, texture, and practical styling moves for today.",
			ReadTime:    "2 min read",
			Tags:        tags[i],
			FetchedAt:   now.Add(-time.Duration(i) * time.Hour),
			ContentHash: id,
		})
	}
	return out
}

func seedLooks(now time.Time, today string) []models.Look {
	names := []string{"Relaxed Neutrals", "City Linen", "Weekend Monochrome", "Soft Utility"}
	tags := [][]string{{"Casual", "Summer"}, {"Linen", "Warm Weather"}, {"Minimal", "Street"}, {"Utility", "Layers"}}
	out := make([]models.Look, 0, 4)
	for i, name := range names {
		id := fmt.Sprintf("look-%d", i+1)
		out = append(out, models.Look{
			ID:             id,
			Source:         "Editorial",
			Title:          name,
			ImageURL:       "https://picsum.photos/seed/" + id + "/700/900",
			SourceURL:      "https://example.com/looks/" + id,
			Tags:           tags[i],
			Season:         "Summer",
			FetchedAt:      now.Add(-time.Duration(i) * time.Hour),
			DisplayDate:    today,
			SelectedForDay: true,
		})
	}
	return out
}

func seedItems(now time.Time, today string) []models.Item {
	brands := []string{"GH Bass", "Morjas", "Sebago", "Aurlands", "Velasca", "Blackstock & Weber"}
	out := make([]models.Item, 0, 6)
	for i, brand := range brands {
		id := fmt.Sprintf("item-%d", i+1)
		out = append(out, models.Item{
			ID:             id,
			Source:         "Retail",
			Brand:          brand,
			Name:           fmt.Sprintf("Polished Loafer %d", i+1),
			Category:       "loafers",
			Price:          fmt.Sprintf("$%d", 160+i*25),
			Currency:       "USD",
			ImageURL:       "https://picsum.photos/seed/" + id + "/700/700",
			ProductURL:     "https://example.com/items/" + id,
			Tags:           []string{"Loafers", "Daily Focus"},
			FetchedAt:      now.Add(-time.Duration(i) * time.Hour),
			DisplayDate:    today,
			SelectedForDay: true,
		})
	}
	return out
}
