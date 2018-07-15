package safari

import (
	"golang.org/x/net/html"
	"io"
)

func GetHighlights(body io.Reader) ([]string, error) {
	highlights := []string{}
	toPrint := false
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return highlights, nil
		case html.TextToken:
			if toPrint {
				highlights = append(highlights, string(z.Text()))
				toPrint = false
			}
		case html.StartTagToken, html.EndTagToken:
			_, value, _ := z.TagAttr()
			if string(value) == "t-annotation-quote annotation-quote" {
				toPrint = true
			}
		}
	}
}
