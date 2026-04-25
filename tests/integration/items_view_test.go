package integration

import (
	"strings"
	"testing"
)

func TestItemsView(t *testing.T) {
	body := get(t, "/items")
	if !strings.Contains(body, "ITEMS") || !strings.Contains(body, "Polished Loafer") {
		t.Fatalf("items view missing expected content")
	}
}
