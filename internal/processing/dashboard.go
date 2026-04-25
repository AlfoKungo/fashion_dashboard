package processing

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/repository"
)

type DashboardService struct {
	store repository.Store
	now   func() time.Time
}

type DashboardData struct {
	ActivePage string
	Today      string
	DateLabel  string
	Quote      string
	Weather    string
	Category   string
	Articles   []models.Article
	Looks      []models.Look
	Items      []models.Item
}

func NewDashboardService(store repository.Store) *DashboardService {
	return &DashboardService{store: store, now: time.Now}
}

func (s *DashboardService) Today(ctx context.Context) (DashboardData, error) {
	now := s.now()
	date := now.Format("2006-01-02")
	articles, err := s.store.ListArticles(ctx, 4)
	if err != nil {
		return DashboardData{}, err
	}
	looks, err := s.store.ListDailyLooks(ctx, date, 4)
	if err != nil {
		return DashboardData{}, err
	}
	items, err := s.store.ListDailyItems(ctx, date, 6)
	if err != nil {
		return DashboardData{}, err
	}
	category := DailyCategory(date)
	if len(items) > 0 {
		category = items[0].Category
	}
	return DashboardData{
		ActivePage: "today",
		Today:      date,
		DateLabel:  now.Format("Monday, January 2"),
		Quote:      "Style is the habit of choosing well.",
		Weather:    "Mild, 72F",
		Category:   category,
		Articles:   CompleteArticles(articles),
		Looks:      CompleteLooks(looks),
		Items:      CompleteItems(items),
	}, nil
}

func (s *DashboardService) Articles(ctx context.Context, amount int) ([]models.Article, error) {
	articles, err := s.store.ListArticles(ctx, amount)
	if err != nil {
		return nil, err
	}
	return CompleteArticles(articles), nil
}

func (s *DashboardService) Looks(ctx context.Context, amount int) ([]models.Look, error) {
	looks, err := s.store.ListLooks(ctx, amount)
	if err != nil {
		return nil, err
	}
	return CompleteLooks(looks), nil
}

func (s *DashboardService) Items(ctx context.Context, amount int) ([]models.Item, error) {
	items, err := s.store.ListItems(ctx, amount)
	if err != nil {
		return nil, err
	}
	return CompleteItems(items), nil
}

func (s *DashboardService) Store() repository.Store {
	return s.store
}
