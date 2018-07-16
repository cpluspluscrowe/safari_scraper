package highlightDb

import (
	"testing"
)

func TestInsert(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
}

func TestGetHighlights(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
	highlights := GetHighlights()
	if len(highlights) <= 0 {
		t.Errorf("GetHighlights did not return any highlights!")
	}
}
