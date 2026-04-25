package unit

import (
	"testing"

	"fashion_dashboard/internal/processing"
)

func TestDailyCategoryDeterministic(t *testing.T) {
	first := processing.DailyCategory("2026-04-25")
	second := processing.DailyCategory("2026-04-25")
	if first != second {
		t.Fatalf("category not deterministic: %q != %q", first, second)
	}
}
