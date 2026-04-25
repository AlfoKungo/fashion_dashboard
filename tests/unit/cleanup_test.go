package unit

import (
	"testing"
	"time"

	"fashion_dashboard/internal/processing"
)

func TestCleanupCutoff(t *testing.T) {
	now := time.Date(2026, 4, 25, 12, 0, 0, 0, time.UTC)
	got := processing.CleanupCutoff(now)
	want := time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("cutoff = %s, want %s", got, want)
	}
}
