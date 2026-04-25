package contract

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestLooksAPI(t *testing.T) {
	app, err := web.NewServer(repository.NewMemoryStore())
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/looks?amount=4", nil)
	app.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var body struct{ Looks []map[string]any }
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Looks) != 4 {
		t.Fatalf("looks = %d", len(body.Looks))
	}
}
