package articles

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestLiveFetcherParsesRSSSource(t *testing.T) {
	feed := `<?xml version="1.0"?>
<rss version="2.0" xmlns:media="http://search.yahoo.com/mrss/" xmlns:dc="http://purl.org/dc/elements/1.1/">
  <channel>
    <item>
      <title><![CDATA[Real Article]]></title>
      <link>https://example.com/article</link>
      <pubDate>Fri, 24 Apr 2026 19:53:55 +0000</pubDate>
      <description><![CDATA[<p>A real article summary from the feed.</p>]]></description>
      <category>Style</category>
      <dc:creator>Editor</dc:creator>
      <media:thumbnail url="https://example.com/image.jpg"/>
    </item>
  </channel>
</rss>`

	fetcher := LiveFetcher{
		Client:  &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) { return xmlResponse(feed), nil })},
		Sources: []Source{{Name: "TEST", URL: "https://example.com/feed"}},
		Now:     func() time.Time { return time.Date(2026, 4, 25, 0, 0, 0, 0, time.UTC) },
	}
	articles, err := fetcher.Fetch(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(articles) != 1 {
		t.Fatalf("articles = %d", len(articles))
	}
	got := articles[0]
	if got.Source != "TEST" || got.Title != "Real Article" || got.URL != "https://example.com/article" || got.ImageURL != "https://example.com/image.jpg" || got.Author != "Editor" {
		t.Fatalf("unexpected article: %+v", got)
	}
}

func TestExtractMetaImage(t *testing.T) {
	page := `<html><head><meta property="og:image" content="https://example.com/og.jpg"></head></html>`
	if got := extractMetaImage(page); got != "https://example.com/og.jpg" {
		t.Fatalf("image = %q", got)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func xmlResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/rss+xml"}},
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}
