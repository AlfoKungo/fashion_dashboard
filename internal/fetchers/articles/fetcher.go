package articles

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"fashion_dashboard/internal/models"
)

type Fetcher interface {
	Fetch(context.Context) ([]models.Article, error)
}

type Source struct {
	Name            string
	URL             string
	AllowCategories []string
}

type LiveFetcher struct {
	Client  *http.Client
	Sources []Source
	Now     func() time.Time
}

func NewLiveFetcher() LiveFetcher {
	return LiveFetcher{
		Client: &http.Client{Timeout: 20 * time.Second},
		Sources: []Source{
			{Name: "GQ", URL: "https://www.gq.com/feed/rss", AllowCategories: []string{"Style", "Shopping", "Watches", "Grooming"}},
			{Name: "HIGHSNOBIETY", URL: "https://www.highsnobiety.com/feeds/rss"},
			{Name: "ESQUIRE", URL: "https://www.esquire.com/rss/style.xml/"},
		},
		Now: time.Now,
	}
}

func (f LiveFetcher) Fetch(ctx context.Context) ([]models.Article, error) {
	client := f.Client
	if client == nil {
		client = &http.Client{Timeout: 20 * time.Second}
	}
	sources := f.Sources
	if len(sources) == 0 {
		sources = NewLiveFetcher().Sources
	}
	now := time.Now
	if f.Now != nil {
		now = f.Now
	}

	var articles []models.Article
	var errs []error
	seen := map[string]bool{}
	for _, source := range sources {
		fetched, err := fetchSource(ctx, client, source, now())
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", source.Name, err))
			continue
		}
		for _, article := range fetched {
			if article.URL == "" || seen[article.URL] {
				continue
			}
			seen[article.URL] = true
			articles = append(articles, article)
		}
	}
	if len(articles) == 0 && len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return articles, nil
}

func fetchSource(ctx context.Context, client *http.Client, source Source, fetchedAt time.Time) ([]models.Article, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, source.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "fashion-dashboard/0.1 (+https://localhost)")
	req.Header.Set("Accept", "application/rss+xml, application/xml, text/xml;q=0.9, */*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("feed returned HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, err
	}

	var feed rssFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}
	out := make([]models.Article, 0, len(feed.Channel.Items))
	for _, item := range feed.Channel.Items {
		article := item.article(source.Name, fetchedAt)
		if article.Title == "" || article.URL == "" {
			continue
		}
		if !source.allows(article.Tags) {
			continue
		}
		if article.ImageURL == "" {
			article.ImageURL = fetchOpenGraphImage(ctx, client, article.URL)
		}
		out = append(out, article)
	}
	return out, nil
}

func (s Source) allows(tags []string) bool {
	if len(s.AllowCategories) == 0 {
		return true
	}
	allowed := map[string]bool{}
	for _, category := range s.AllowCategories {
		allowed[strings.ToLower(category)] = true
	}
	for _, tag := range tags {
		if allowed[strings.ToLower(tag)] {
			return true
		}
	}
	return false
}

func fetchOpenGraphImage(ctx context.Context, client *http.Client, pageURL string) string {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("User-Agent", "fashion-dashboard/0.1 (+https://localhost)")
	req.Header.Set("Accept", "text/html, */*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return ""
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return ""
	}
	return extractMetaImage(string(body))
}

type rssFeed struct {
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Items []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	GUID        string      `xml:"guid"`
	PubDate     string      `xml:"pubDate"`
	Description string      `xml:"description"`
	Creator     string      `xml:"creator"`
	Categories  []string    `xml:"category"`
	Media       []mediaNode `xml:"content"`
	Thumbnails  []mediaNode `xml:"thumbnail"`
	Enclosures  []mediaNode `xml:"enclosure"`
}

type mediaNode struct {
	URL string `xml:"url,attr"`
}

func (i rssItem) article(source string, fetchedAt time.Time) models.Article {
	link := strings.TrimSpace(i.Link)
	if link == "" {
		link = strings.TrimSpace(i.GUID)
	}
	summary := cleanText(i.Description)
	publishedAt := parseRSSDate(i.PubDate)
	return models.Article{
		Source:      source,
		Title:       cleanText(i.Title),
		URL:         link,
		ImageURL:    i.imageURL(),
		Author:      cleanText(i.Creator),
		PublishedAt: publishedAt,
		Summary:     summarize(summary),
		ReadTime:    estimateReadTime(summary),
		Tags:        cleanTags(i.Categories),
		FetchedAt:   fetchedAt,
	}
}

func (i rssItem) imageURL() string {
	for _, candidates := range [][]mediaNode{i.Thumbnails, i.Media, i.Enclosures} {
		for _, candidate := range candidates {
			if candidate.URL != "" {
				return candidate.URL
			}
		}
	}
	return ""
}

func parseRSSDate(value string) time.Time {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}
	}
	layouts := []string{time.RFC1123Z, time.RFC1123, time.RFC822Z, time.RFC822, time.RFC3339}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t
		}
	}
	return time.Time{}
}

var tagRE = regexp.MustCompile(`<[^>]+>`)
var spaceRE = regexp.MustCompile(`\s+`)
var metaRE = regexp.MustCompile(`(?is)<meta\s+[^>]*>`)
var propertyRE = regexp.MustCompile(`(?is)\s(?:property|name)=["']([^"']+)["']`)
var contentRE = regexp.MustCompile(`(?is)\scontent=["']([^"']+)["']`)

func cleanText(value string) string {
	value = html.UnescapeString(value)
	value = tagRE.ReplaceAllString(value, " ")
	value = spaceRE.ReplaceAllString(value, " ")
	return strings.TrimSpace(value)
}

func summarize(value string) string {
	value = cleanText(value)
	if value == "" {
		return "Latest men's fashion and style coverage from the source."
	}
	runes := []rune(value)
	if len(runes) <= 180 {
		return value
	}
	return strings.TrimSpace(string(runes[:177])) + "..."
}

func estimateReadTime(value string) string {
	words := len(strings.Fields(cleanText(value)))
	minutes := words / 225
	if minutes < 1 {
		minutes = 2
	}
	return fmt.Sprintf("%d min read", minutes)
}

func cleanTags(tags []string) []string {
	out := make([]string, 0, len(tags))
	seen := map[string]bool{}
	for _, tag := range tags {
		tag = cleanText(tag)
		if tag == "" || seen[tag] {
			continue
		}
		seen[tag] = true
		out = append(out, tag)
	}
	if len(out) == 0 {
		out = append(out, "Style")
	}
	return out
}

func extractMetaImage(page string) string {
	for _, meta := range metaRE.FindAllString(page, -1) {
		prop := attr(propertyRE, meta)
		if prop != "og:image" && prop != "twitter:image" {
			continue
		}
		if content := attr(contentRE, meta); content != "" {
			return html.UnescapeString(content)
		}
	}
	return ""
}

func attr(re *regexp.Regexp, value string) string {
	matches := re.FindStringSubmatch(value)
	if len(matches) < 2 {
		return ""
	}
	return strings.TrimSpace(matches[1])
}
