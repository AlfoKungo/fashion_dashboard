package scheduler

import (
	"context"
	"time"

	articlefetcher "fashion_dashboard/internal/fetchers/articles"
	itemfetcher "fashion_dashboard/internal/fetchers/items"
	lookfetcher "fashion_dashboard/internal/fetchers/looks"
	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/processing"
	"fashion_dashboard/internal/repository"
)

type Workflow struct {
	store    repository.Store
	articles articlefetcher.Fetcher
	looks    lookfetcher.Fetcher
	items    itemfetcher.Fetcher
	now      func() time.Time
}

func NewWorkflow(store repository.Store) *Workflow {
	return NewWorkflowWithFetchers(store, articlefetcher.NewLiveFetcher(), lookfetcher.SampleFetcher{}, itemfetcher.SampleFetcher{})
}

func NewWorkflowWithFetchers(store repository.Store, articles articlefetcher.Fetcher, looks lookfetcher.Fetcher, items itemfetcher.Fetcher) *Workflow {
	return &Workflow{
		store:    store,
		articles: articles,
		looks:    looks,
		items:    items,
		now:      time.Now,
	}
}

func (w *Workflow) Run(ctx context.Context) error {
	now := w.now()
	date := now.Format("2006-01-02")
	category := processing.DailyCategory(date)

	articles, err := w.articles.Fetch(ctx)
	if err != nil {
		return err
	}
	looks, err := w.looks.Fetch(ctx)
	if err != nil {
		return err
	}
	items, err := w.items.Fetch(ctx, category)
	if err != nil {
		return err
	}

	articles = processing.NormalizeArticles(articles, now)
	looks = processing.NormalizeLooks(looks, now)
	items = processing.NormalizeItems(items, now)

	if err := w.store.UpsertArticles(ctx, articles); err != nil {
		return err
	}
	if err := w.store.UpsertLooks(ctx, looks); err != nil {
		return err
	}
	if err := w.store.UpsertItems(ctx, items); err != nil {
		return err
	}

	selectedLooks := processing.SelectLooksForDay(looks, date, 4)
	selectedItems := processing.SelectItemsForDay(items, date, category, 6)
	if err := w.store.MarkDailyLooks(ctx, date, idsOfLooks(selectedLooks)); err != nil {
		return err
	}
	if err := w.store.MarkDailyItems(ctx, date, idsOfItems(selectedItems)); err != nil {
		return err
	}
	if err := w.store.SaveTrendSummary(ctx, processing.GenerateTrendSummary(date, articles, selectedLooks, selectedItems)); err != nil {
		return err
	}
	return processing.CleanupOldData(ctx, w.store, now)
}

func idsOfLooks(looks []models.Look) []string {
	ids := make([]string, 0, len(looks))
	for _, look := range looks {
		ids = append(ids, look.ID)
	}
	return ids
}

func idsOfItems(items []models.Item) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}
