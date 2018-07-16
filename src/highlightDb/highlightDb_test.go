package highlightDb

import (
	"fmt"
	"testing"
)

func TestInsert(*testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
}

func TestGetHighlights(*testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
	highlights := GetHighlights()
	fmt.Println(highlights)
}
