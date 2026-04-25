package items

import (
	"context"
	"fmt"
	"time"

	"fashion_dashboard/internal/models"
)

type Fetcher interface {
	Fetch(context.Context, string) ([]models.Item, error)
}

type SampleFetcher struct{}

func (SampleFetcher) Fetch(_ context.Context, category string) ([]models.Item, error) {
	now := time.Now()
	out := make([]models.Item, 0, 6)
	for i := 1; i <= 6; i++ {
		id := fmt.Sprintf("fetched-item-%d", i)
		out = append(out, models.Item{
			ID:         id,
			Source:     "Retail",
			Brand:      fmt.Sprintf("Brand %d", i),
			Name:       fmt.Sprintf("%s pick %d", category, i),
			Category:   category,
			Price:      fmt.Sprintf("$%d", 120+i*15),
			Currency:   "USD",
			ImageURL:   "https://picsum.photos/seed/" + id + "/700/700",
			ProductURL: "https://example.com/fetched/items/" + id,
			Tags:       []string{category},
			FetchedAt:  now,
		})
	}
	return out, nil
}
