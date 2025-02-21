package main

import (
	"fmt"
	"os"
	"strconv"

	gotwitter "github.com/928799934/go-twitter"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
	BearerToken       = ""
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for sampling count.")
		os.Exit(1)
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	client := newOAuth2Client()

	samplingTweets(client, count)
}

func newOAuth2Client() *gotwitter.GoTwitter {

	if BearerToken != "" {
		return gotwitter.NewGoTwitter(gotwitter.WithBearerToken(BearerToken))
	}

	return gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))
}
