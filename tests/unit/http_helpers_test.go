package unit

import (
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/web"
)

func TestParseAmount(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/articles", nil)
	got, err := web.ParseAmount(req, 4)
	if err != nil || got != 4 {
		t.Fatalf("default amount = %d, %v", got, err)
	}

	req = httptest.NewRequest("GET", "/api/articles?amount=50", nil)
	got, err = web.ParseAmount(req, 4)
	if err != nil || got != 50 {
		t.Fatalf("amount = %d, %v", got, err)
	}

	for _, path := range []string{"/api/articles?amount=0", "/api/articles?amount=51", "/api/articles?amount=abc"} {
		req = httptest.NewRequest("GET", path, nil)
		if _, err := web.ParseAmount(req, 4); err == nil {
			t.Fatalf("expected error for %s", path)
		}
	}
}
