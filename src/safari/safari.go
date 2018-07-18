package safari

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
)

type safariHighlight struct {
	Text   string
	Source string
}

func GetSafariHighlights() []string {
	highlights := getSafariHighlights()
	withCitations := addCitationToHighlights(highlights)
	return withCitations
}

func getSafariHighlights() []safariHighlight {
	url := "https://www.safaribooksonline.com/u/8aaaf02c-85f9-42f8-8628-3e0c83d24fd6/"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	highlights, err := extractHighlights(resp.Body)
	if err != nil {
		panic(err)
	}
	return highlights
}

func addCitationToHighlights(safariHighlights []safariHighlight) []string {
	highlights := []string{}
	for _, highlight := range safariHighlights {
		highlights = append(highlights, highlight.Text+" - \""+highlight.Source+"\"")
	}
	return highlights
}

func extractHighlights(body io.Reader) ([]safariHighlight, error) {
	highlights := []safariHighlight{}
	toPrint := false
	z := html.NewTokenizer(body)
	toGetSource := false
	source := ""
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return highlights, nil
		case html.TextToken:
			if toGetSource {
				source = string(z.Text())
				toGetSource = false
			}
			if toPrint {
				highlights = append(highlights, safariHighlight{Text: string(z.Text()), Source: source})
				toPrint = false
			}
		case html.StartTagToken, html.EndTagToken:
			_, value, _ := z.TagAttr()
			if string(value) == "t-annotation-quote annotation-quote" {
				toPrint = true
			}
			if string(value) == "t-annotation-archive-title" {
				toGetSource = true
			}
		}
	}
}

//
