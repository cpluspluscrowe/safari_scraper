package main

import (
	"fmt"
	"highlightDb"
	"safari"
	"time"
	"twitter"
)

func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(120000*time.Millisecond, postHighlightsToTwitter)
}

func postHighlightsToTwitter() {
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
	highlightDb.InsertHighlights(highlights)

	dbHighlights := highlightDb.GetUnpostedHighlights()
	fmt.Printf("Number of highlights to post: %d\n", len(dbHighlights))
	for _, highlight := range dbHighlights {
		twitter.Tweet(highlight.Text)
		highlightDb.SetHighlightAsPosted(highlight.Text)
	}
}
