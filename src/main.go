package main

import (
	"fmt"
	"net/http"
	"safari"
	"twitter"
)

func main() {
	url := "https://www.safaribooksonline.com/u/8aaaf02c-85f9-42f8-8628-3e0c83d24fd6/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	highlights, err := safari.GetHighlights(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(highlights)
	for _, highlight := range highlights {
		twitter.Tweet(highlight)
	}
}
