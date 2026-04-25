package processing

import (
	"context"
	"time"

	"fashion_dashboard/internal/repository"
)

func CleanupOldData(ctx context.Context, store repository.Store, now time.Time) error {
	cutoff := CleanupCutoff(now)
	if err := store.DeleteArticlesOlderThan(ctx, cutoff); err != nil {
		return err
	}
	if err := store.DeleteLooksOlderThan(ctx, cutoff); err != nil {
		return err
	}
	if err := store.DeleteItemsOlderThan(ctx, cutoff); err != nil {
		return err
	}
	return store.DeleteTrendSummariesOlderThan(ctx, cutoff)
}
