package integration

import (
	"strings"
	"testing"
)

func TestLooksView(t *testing.T) {
	body := get(t, "/looks")
	if !strings.Contains(body, "LOOKS") || !strings.Contains(body, `href="/looks">LOOKS`) {
		t.Fatalf("looks view missing expected content")
	}
}
