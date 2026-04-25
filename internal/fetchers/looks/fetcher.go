package looks

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
)

type Fetcher interface {
	Fetch(context.Context) ([]models.Look, error)
}

type SampleFetcher struct{}

func (SampleFetcher) Fetch(_ context.Context) ([]models.Look, error) {
	now := time.Now()
	return []models.Look{
		{ID: "fetched-look-1", Source: "Editorial", Title: "Light Utility", ImageURL: "https://picsum.photos/seed/fetched-look-1/700/900", SourceURL: "https://example.com/fetched/looks/1", Tags: []string{"Utility", "Casual"}, Season: "Summer", FetchedAt: now},
		{ID: "fetched-look-2", Source: "Editorial", Title: "Evening Neutrals", ImageURL: "https://picsum.photos/seed/fetched-look-2/700/900", SourceURL: "https://example.com/fetched/looks/2", Tags: []string{"Neutral"}, Season: "Summer", FetchedAt: now},
		{ID: "fetched-look-3", Source: "Editorial", Title: "Weekend Linen", ImageURL: "https://picsum.photos/seed/fetched-look-3/700/900", SourceURL: "https://example.com/fetched/looks/3", Tags: []string{"Linen"}, Season: "Summer", FetchedAt: now},
		{ID: "fetched-look-4", Source: "Editorial", Title: "Soft Tailoring", ImageURL: "https://picsum.photos/seed/fetched-look-4/700/900", SourceURL: "https://example.com/fetched/looks/4", Tags: []string{"Tailoring"}, Season: "Summer", FetchedAt: now},
	}, nil
}
