package contract

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestArticlesAPI(t *testing.T) {
	app, err := web.NewServer(repository.NewMemoryStore())
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/articles?amount=4", nil)
	app.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var body struct{ Articles []map[string]any }
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Articles) != 4 {
		t.Fatalf("articles = %d", len(body.Articles))
	}
}
