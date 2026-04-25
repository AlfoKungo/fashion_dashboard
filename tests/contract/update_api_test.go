package contract

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/web"
)

func TestUpdateAPI(t *testing.T) {
	app, err := web.NewServerWithWorkflow(repository.NewMemoryStore(), fakeRunner{})
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/update", nil)
	app.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", rec.Code, rec.Body.String())
	}
	var body struct {
		Status   string
		Date     string
		Category string
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Status != "updated" || body.Date == "" || body.Category == "" {
		t.Fatalf("unexpected response: %+v", body)
	}
}

type fakeRunner struct{}

func (fakeRunner) Run(context.Context) error { return nil }
