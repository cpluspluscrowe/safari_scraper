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
	parseBody(resp.Body)
	fmt.Println("Done")
}

func parseBody(body io.Reader) {
	doc, err := html.Parse(body)
	if err != nil {
		return
	}
	recurse(doc)
}

func recurse(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		recurse(c)
	}
}
