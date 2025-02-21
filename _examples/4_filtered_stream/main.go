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
	BearerToken       = ""
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("The 1st parameter for command is required. (create|stream)")
		os.Exit(1)
	}

	command := args[1]

	switch command {
	case "list":
		// list search stream rules
		listSearchStreamRules()
	case "delete":
		// delete a specified rule
		if len(args) < 3 {
			fmt.Println("The 2nd parameter for rule ID to delete is required.")
			os.Exit(1)
		}

		ruleID := args[2]
		deleteSearchStreamRules(ruleID)
	case "create":
		// create a search stream rule
		if len(args) < 3 {
			fmt.Println("The 2nd parameter for keyword of search stream rule is required.")
			os.Exit(1)
		}

		keyword := args[2]
		createSearchStreamRules(keyword)
	case "stream":
		// exec filtered stream API
		execSearchStream()
	default:
		fmt.Println("Undefined command. Command should be 'create' or 'stream'.")
		os.Exit(1)
	}
}

func newOAuth2Client() *gotwitter.GoTwitter {
	if BearerToken != "" {
		return gotwitter.NewGoTwitter(gotwitter.WithBearerToken(BearerToken))
	}
	return gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))
}
