package repository

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
)

type ArticleRepository interface {
	ListArticles(ctx context.Context, amount int) ([]models.Article, error)
	UpsertArticles(ctx context.Context, articles []models.Article) error
	DeleteArticlesOlderThan(ctx context.Context, cutoff time.Time) error
}

type LookRepository interface {
	ListLooks(ctx context.Context, amount int) ([]models.Look, error)
	ListDailyLooks(ctx context.Context, date string, amount int) ([]models.Look, error)
	UpsertLooks(ctx context.Context, looks []models.Look) error
	MarkDailyLooks(ctx context.Context, date string, ids []string) error
	DeleteLooksOlderThan(ctx context.Context, cutoff time.Time) error
}

type ItemRepository interface {
	ListItems(ctx context.Context, amount int) ([]models.Item, error)
	ListDailyItems(ctx context.Context, date string, amount int) ([]models.Item, error)
	UpsertItems(ctx context.Context, items []models.Item) error
	MarkDailyItems(ctx context.Context, date string, ids []string) error
	DeleteItemsOlderThan(ctx context.Context, cutoff time.Time) error
}

type ImageRepository interface {
	GetArticleImage(ctx context.Context, id string) (models.Image, bool, error)
	GetLookImage(ctx context.Context, id string) (models.Image, bool, error)
	GetItemImage(ctx context.Context, id string) (models.Image, bool, error)
}

type TrendRepository interface {
	SaveTrendSummary(ctx context.Context, summary models.TrendSummary) error
	DeleteTrendSummariesOlderThan(ctx context.Context, cutoff time.Time) error
}

type Store interface {
	ArticleRepository
	LookRepository
	ItemRepository
	ImageRepository
	TrendRepository
}
