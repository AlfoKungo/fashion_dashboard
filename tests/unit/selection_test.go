package unit

import (
	"fmt"
	"testing"

	"fashion_dashboard/internal/models"
	"fashion_dashboard/internal/processing"
)

func TestSelectionCounts(t *testing.T) {
	var looks []models.Look
	var items []models.Item
	for i := 0; i < 8; i++ {
		looks = append(looks, look(fmt.Sprintf("look-%d", i)))
		items = append(items, item(fmt.Sprintf("item-%d", i), "loafers"))
	}
	if got := processing.SelectLooksForDay(looks, "2026-04-25", 4); len(got) != 4 {
		t.Fatalf("looks = %d", len(got))
	}
	if got := processing.SelectItemsForDay(items, "2026-04-25", "loafers", 6); len(got) != 6 {
		t.Fatalf("items = %d", len(got))
	}
}
