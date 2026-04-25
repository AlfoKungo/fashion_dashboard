package processing

import (
	"fmt"
	"time"

	"fashion_dashboard/internal/models"
)

func GenerateTrendSummary(date string, articles []models.Article, looks []models.Look, items []models.Item) models.TrendSummary {
	category := "daily style"
	if len(items) > 0 {
		category = items[0].Category
	}
	return models.TrendSummary{
		ID:        HashCanonical(date),
		Date:      date,
		Summary:   fmt.Sprintf("Today's edit combines %d articles, %d looks, and a focused %s item selection.", len(articles), len(looks), category),
		CreatedAt: time.Now(),
	}
}
