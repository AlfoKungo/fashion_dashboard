package repository

import (
	"context"
	"sync"
	"time"

	"fashion_dashboard/internal/models"
)

type MemoryStore struct {
	mu       sync.RWMutex
	articles []models.Article
	looks    []models.Look
	items    []models.Item
	trends   []models.TrendSummary
}

func NewMemoryStore() *MemoryStore {
	now := time.Now()
	today := now.Format("2006-01-02")
	return &MemoryStore{
		articles: seedArticles(now),
		looks:    seedLooks(now, today),
		items:    seedItems(now, today),
		trends:   []models.TrendSummary{{ID: "trend-1", Date: today, Summary: "Relaxed tailoring, breathable layers, and polished footwear define today's edit.", CreatedAt: now}},
	}
}

func (s *MemoryStore) ListArticles(_ context.Context, amount int) ([]models.Article, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	articles := append([]models.Article(nil), s.articles...)
	articles = orderArticlesByFreshnessAndSource(articles)
	limit := min(amount, len(articles))
	for i := range articles[:limit] {
		articles[i].ImageSrc = "/images/articles/" + articles[i].ID
	}
	return articles[:limit], nil
}

func (s *MemoryStore) UpsertArticles(_ context.Context, articles []models.Article) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, article := range articles {
		if article.ID == "" {
			article.ID = article.ContentHash
		}
		replaced := false
		for i := range s.articles {
			if s.articles[i].URL == article.URL {
				s.articles[i] = article
				replaced = true
				break
			}
		}
		if !replaced {
			s.articles = append(s.articles, article)
		}
	}
	return nil
}

func (s *MemoryStore) DeleteArticlesOlderThan(_ context.Context, cutoff time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.articles = filterByTime(s.articles, cutoff, func(a models.Article) time.Time { return a.FetchedAt })
	return nil
}

func (s *MemoryStore) ListLooks(_ context.Context, amount int) ([]models.Look, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	looks := append([]models.Look(nil), s.looks...)
	limit := min(amount, len(looks))
	for i := range looks[:limit] {
		looks[i].ImageSrc = "/images/looks/" + looks[i].ID
	}
	return looks[:limit], nil
}

func (s *MemoryStore) ListDailyLooks(_ context.Context, date string, amount int) ([]models.Look, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []models.Look
	for _, look := range s.looks {
		if look.DisplayDate == date && look.SelectedForDay {
			look.ImageSrc = "/images/looks/" + look.ID
			out = append(out, look)
		}
	}
	if len(out) == 0 {
		out = append([]models.Look(nil), s.looks...)
		for i := range out {
			out[i].ImageSrc = "/images/looks/" + out[i].ID
		}
	}
	return out[:min(amount, len(out))], nil
}

func (s *MemoryStore) UpsertLooks(_ context.Context, looks []models.Look) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, look := range looks {
		replaced := false
		for i := range s.looks {
			if s.looks[i].SourceURL == look.SourceURL {
				s.looks[i] = look
				replaced = true
				break
			}
		}
		if !replaced {
			s.looks = append(s.looks, look)
		}
	}
	return nil
}

func (s *MemoryStore) MarkDailyLooks(_ context.Context, date string, ids []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	selected := stringSet(ids)
	for i := range s.looks {
		s.looks[i].SelectedForDay = selected[s.looks[i].ID]
		if s.looks[i].SelectedForDay {
			s.looks[i].DisplayDate = date
		}
	}
	return nil
}

func (s *MemoryStore) DeleteLooksOlderThan(_ context.Context, cutoff time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.looks = filterByTime(s.looks, cutoff, func(l models.Look) time.Time { return l.FetchedAt })
	return nil
}

func (s *MemoryStore) ListItems(_ context.Context, amount int) ([]models.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := append([]models.Item(nil), s.items...)
	limit := min(amount, len(items))
	for i := range items[:limit] {
		items[i].ImageSrc = "/images/items/" + items[i].ID
	}
	return items[:limit], nil
}

func (s *MemoryStore) ListDailyItems(_ context.Context, date string, amount int) ([]models.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []models.Item
	for _, item := range s.items {
		if item.DisplayDate == date && item.SelectedForDay {
			item.ImageSrc = "/images/items/" + item.ID
			out = append(out, item)
		}
	}
	if len(out) == 0 {
		out = append([]models.Item(nil), s.items...)
		for i := range out {
			out[i].ImageSrc = "/images/items/" + out[i].ID
		}
	}
	return out[:min(amount, len(out))], nil
}

func (s *MemoryStore) UpsertItems(_ context.Context, items []models.Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range items {
		replaced := false
		for i := range s.items {
			if s.items[i].ProductURL == item.ProductURL {
				s.items[i] = item
				replaced = true
				break
			}
		}
		if !replaced {
			s.items = append(s.items, item)
		}
	}
	return nil
}

func (s *MemoryStore) MarkDailyItems(_ context.Context, date string, ids []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	selected := stringSet(ids)
	for i := range s.items {
		s.items[i].SelectedForDay = selected[s.items[i].ID]
		if s.items[i].SelectedForDay {
			s.items[i].DisplayDate = date
		}
	}
	return nil
}

func (s *MemoryStore) DeleteItemsOlderThan(_ context.Context, cutoff time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = filterByTime(s.items, cutoff, func(i models.Item) time.Time { return i.FetchedAt })
	return nil
}

func (s *MemoryStore) SaveTrendSummary(_ context.Context, summary models.TrendSummary) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.trends {
		if s.trends[i].Date == summary.Date {
			s.trends[i] = summary
			return nil
		}
	}
	s.trends = append(s.trends, summary)
	return nil
}

func (s *MemoryStore) DeleteTrendSummariesOlderThan(_ context.Context, cutoff time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.trends = filterByTime(s.trends, cutoff, func(t models.TrendSummary) time.Time { return t.CreatedAt })
	return nil
}

func filterByTime[T any](values []T, cutoff time.Time, at func(T) time.Time) []T {
	out := values[:0]
	for _, value := range values {
		if !at(value).Before(cutoff) {
			out = append(out, value)
		}
	}
	return out
}

func stringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}
