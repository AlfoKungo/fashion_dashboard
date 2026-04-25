package repository

import (
	"context"

	"fashion_dashboard/internal/models"
)

func (s *MemoryStore) GetArticleImage(_ context.Context, id string) (models.Image, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, article := range s.articles {
		if article.ID == id {
			return models.Image{ID: id, URL: article.ImageURL, Bytes: article.ImageBytes, ContentType: article.ImageContentType}, true, nil
		}
	}
	return models.Image{}, false, nil
}

func (s *MemoryStore) GetLookImage(_ context.Context, id string) (models.Image, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, look := range s.looks {
		if look.ID == id {
			return models.Image{ID: id, URL: look.ImageURL, Bytes: look.ImageBytes, ContentType: look.ImageContentType}, true, nil
		}
	}
	return models.Image{}, false, nil
}

func (s *MemoryStore) GetItemImage(_ context.Context, id string) (models.Image, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, item := range s.items {
		if item.ID == id {
			return models.Image{ID: id, URL: item.ImageURL, Bytes: item.ImageBytes, ContentType: item.ImageContentType}, true, nil
		}
	}
	return models.Image{}, false, nil
}
