package integration

import (
	"strings"
	"testing"
)

func TestArticlesView(t *testing.T) {
	body := get(t, "/articles")
	if !strings.Contains(body, "ARTICLES") || !strings.Contains(body, "2 min read") {
		t.Fatalf("articles view missing expected content")
	}
}
