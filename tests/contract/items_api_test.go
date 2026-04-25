package contract

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestItemsAPI(t *testing.T) {
	app, err := web.NewServer(repository.NewMemoryStore())
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/items?amount=6", nil)
	app.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var body struct {
		Category string
		Items    []map[string]any
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Category == "" || len(body.Items) != 6 {
		t.Fatalf("category=%q items=%d", body.Category, len(body.Items))
	}
}
