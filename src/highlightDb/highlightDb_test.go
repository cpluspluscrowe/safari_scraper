package highlightDb

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
}

func TestGetHighlights(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
	highlights := GetUnpostedHighlights()
	if len(highlights) <= 0 {
		t.Errorf("GetHighlights did not return any highlights!")
	} else {
		for _, highlight := range highlights {
			fmt.Println(highlight.Text)
		}
	}
}
