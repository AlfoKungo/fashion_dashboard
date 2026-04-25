package contract

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestImageRoutes(t *testing.T) {
	app, err := web.NewServer(repository.NewMemoryStore())
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"/images/articles/article-1", "/images/looks/look-1", "/images/items/item-1"} {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		app.Handler().ServeHTTP(rec, req)
		if rec.Code != http.StatusFound && rec.Code != http.StatusOK {
			t.Fatalf("%s status = %d", path, rec.Code)
		}
	}
}
