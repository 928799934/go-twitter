package main

import (
	"fmt"
	"os"

	gotwitter "github.com/928799934/go-twitter"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for command. (post|delete)")
		os.Exit(1)
	}
	cmd := args[1]

	if len(args) < 3 {
		fmt.Println("The 1st parameter is required. (tweet_text|tweet_id)")
		os.Exit(1)
	}
	p := args[2]

	oauth1Client := newOAuth1Client()

	switch cmd {
	case "post":
		tweetID, err := SimpleTweet(oauth1Client, p)
		if err != nil {
			panic(err)
		}

		fmt.Println("Posted tweet ID is ", tweetID)
	case "delete":
		b, err := DeleteTweet(oauth1Client, p)
		if err != nil {
			panic(err)
		}

		fmt.Println("Delete tweet result:  ", b)
	default:
		fmt.Println("Unsupported command. Supported commands are 'post' and 'delete'.")
	}
}

func newOAuth1Client() *gotwitter.GoTwitter {
	return gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))
}
