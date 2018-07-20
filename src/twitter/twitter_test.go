package twitter

import (
	"testing"
)

func TestGetTweets(t *testing.T) {
	GetTweets("other")

}

func TestSearch(t *testing.T) {
	Search("golang")
}
