package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

func main() {
	url := "https://www.safaribooksonline.com/u/8aaaf02c-85f9-42f8-8628-3e0c83d24fd6/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	highlights, err := parseBody(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(highlights)
}

func parseBody(body io.Reader) ([]string, error) {
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
