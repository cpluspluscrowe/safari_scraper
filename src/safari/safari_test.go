package safari

import (
	"testing"
)

func TestItRetrievesHighlights(t *testing.T) {
	safariHighlights := GetSafariHighlights()
	if len(safariHighlights) == 0 {
		t.Errorf("GetHighlights did not retrieve any safari highlights; received: %d", len(safariHighlights))
	}
}

func TestThatHighlightsHaveSource(t *testing.T) {
	safariHighlights := GetSafariHighlights()
	allHighlightsWithEmptySource := true
	if len(safariHighlights) > 0 {
		for _, highlight := range safariHighlights {
			if highlight.Source != "" {
				allHighlightsWithEmptySource = false
			}
		}
	}
	if allHighlightsWithEmptySource == true {
		t.Errorf("Sources should not be empty string")
	}
}
