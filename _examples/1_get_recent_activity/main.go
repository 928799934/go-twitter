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
		fmt.Println("The 1st parameter is required for your account id.")
		os.Exit(1)
	}

	accountID := args[1]

	oauth1Client := newOAuth1Client()
	onlyFollowsRecentActivity(oauth1Client, accountID)
}

func newOAuth1Client() *gotwitter.GoTwitter {
	return gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))
}
