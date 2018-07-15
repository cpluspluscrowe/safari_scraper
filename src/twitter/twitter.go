package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Tweet(tweetText string) {
	config := oauth1.NewConfig("GrH8fCv9zuecdPlvcfrPJVUqE", "7QHnRRRktOpa4EtnQGAAenMSqCgbmYZi2m5Udy9gVUc8px0gOi")
	token := oauth1.NewToken("3127629293-7WR9PxYovLnrOZytacX6IGujOdQKUt295mnz3GM", "JKF32Yv8eAZjdubrBTR9z1gF3BBqrbAKrBweL6AnoRr4B")
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
