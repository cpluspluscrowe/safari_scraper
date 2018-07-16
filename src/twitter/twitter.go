package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Tweet(tweetText string) {
	config := oauth1.NewConfig("wvXNTaqdmmeAj884c0n4fUI1T", "vKQrJC5pO78U7cv38bo7NaisQowPc8RIFYoNgHqFDfmNZy0YWR")
	token := oauth1.NewToken("3127629293-GJ3MaJe9QA540xRgMHFQz9Zo2OWDN5mSU4jc3DB", "PNiLZ8POuJWbCUUdaV4IFELUlz08YcwL7L8uJeXQPNjNu")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	//	tweet, resp, err :=
	tweet, resp, err := client.Statuses.Update(tweetText, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet, resp)
}

func FakeTweet(tweetText string) {
	fmt.Println(tweetText)
}
