package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestTodayView(t *testing.T) {
	body := get(t, "/")
	for _, want := range []string{"TOP ARTICLES", "DAILY INSPIRATION", "DAILY ITEM FOCUS", "active", "MEN'S FASHION DASHBOARD"} {
		if !strings.Contains(body, want) {
			t.Fatalf("today view missing %q", want)
		}
	}
}

func get(t *testing.T, path string) string {
	t.Helper()
	app, err := web.NewServer(repository.NewMemoryStore())
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	app.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("%s status = %d: %s", path, rec.Code, rec.Body.String())
	}
	return rec.Body.String()
}
