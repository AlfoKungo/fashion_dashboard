package processing

import (
	"context"
	"io"
	"net/http"
	"time"

	"fashion_dashboard/internal/models"
)

type DownloadedImage struct {
	URL         string
	Bytes       []byte
	ContentType string
}

func DownloadImage(ctx context.Context, client *http.Client, url string) DownloadedImage {
	if url == "" {
		return DownloadedImage{}
	}
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return DownloadedImage{URL: url}
	}
	resp, err := client.Do(req)
	if err != nil {
		return DownloadedImage{URL: url}
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return DownloadedImage{URL: url}
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 5<<20))
	if err != nil {
		return DownloadedImage{URL: url}
	}
	return DownloadedImage{URL: url, Bytes: body, ContentType: resp.Header.Get("Content-Type")}
}

func CompleteArticles(in []models.Article) []models.Article {
	out := make([]models.Article, 0, len(in))
	for _, article := range in {
		if article.Title == "" || article.Source == "" || article.Summary == "" || article.ReadTime == "" {
			continue
		}
		out = append(out, article)
	}
	return out
}

func CompleteLooks(in []models.Look) []models.Look {
	out := make([]models.Look, 0, len(in))
	for _, look := range in {
		if !completeLook(look) {
			continue
		}
		out = append(out, look)
	}
	return out
}

func CompleteItems(in []models.Item) []models.Item {
	out := make([]models.Item, 0, len(in))
	for _, item := range in {
		if !completeItem(item) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func completeLook(look models.Look) bool {
	return look.Title != "" && look.Source != ""
}

func completeItem(item models.Item) bool {
	return item.Brand != "" && item.Name != "" && item.Category != "" && item.Price != ""
}
