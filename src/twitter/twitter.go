package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
)

func Tweet(tweetText string) {
	httpClient := getClient()

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	//	tweet, resp, err :=
	_, _, err := client.Statuses.Update(tweetText, nil)
	if err != nil {
		panic(err)
	}
}

func FakeTweet(tweetText string) {
	fmt.Println(tweetText)
}

func GetTweets(highlight string) {
	httpClient := getClient()
	client := twitter.NewClient(httpClient)

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: highlight,
	})
	fmt.Println(search, resp, err)
}

func getClient() *http.Client {
	config := oauth1.NewConfig("wvXNTaqdmmeAj884c0n4fUI1T", "vKQrJC5pO78U7cv38bo7NaisQowPc8RIFYoNgHqFDfmNZy0YWR")
	token := oauth1.NewToken("3127629293-GJ3MaJe9QA540xRgMHFQz9Zo2OWDN5mSU4jc3DB", "PNiLZ8POuJWbCUUdaV4IFELUlz08YcwL7L8uJeXQPNjNu")
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}
