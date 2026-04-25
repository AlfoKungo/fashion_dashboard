package integration

import (
	"strings"
	"testing"
)

func TestPartialContentNoEmptyCards(t *testing.T) {
	body := get(t, "/")
	if strings.Contains(body, "<article class=\"card item-card\"></article>") {
		t.Fatalf("rendered empty item card")
	}
}
