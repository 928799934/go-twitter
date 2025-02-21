package main

import (
	"context"
	"fmt"

	gotwitter "github.com/928799934/go-twitter"
	. "github.com/928799934/go-twitter/_examples"
	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/tweet/searchtweet"
	sttypes "github.com/928799934/go-twitter/tweet/searchtweet/types"
	"github.com/928799934/go-twitter/user/follow"
	ftypes "github.com/928799934/go-twitter/user/follow/types"
)

type twitterUser struct {
	ID       string
	Name     string
	Username string
}

func (f twitterUser) displayName() string {
	return fmt.Sprintf("%s@%s", f.Name, f.Username)
}

// onlyFollowsRecentActivity will output the accounts that are unilaterally following
// the specified user ID, along with up to three most recent tweets.
func onlyFollowsRecentActivity(c *gotwitter.GoTwitter, userID string) {
	// list follows
	followings := map[string]twitterUser{}

	paginationToken := "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowingsInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowings(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followings[StringValue(u.ID)] = twitterUser{
				ID:       StringValue(u.ID),
				Name:     StringValue(u.Name),
				Username: StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// list followers
	followers := map[string]twitterUser{}

	paginationToken = "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowersInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowers(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followers[StringValue(u.ID)] = twitterUser{
				ID:       StringValue(u.ID),
				Name:     StringValue(u.Name),
				Username: StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// only following
	onlyFollowings := map[string]twitterUser{}
	for fid, u := range followings {
		if _, ok := followers[fid]; ok {
			continue
		}

		onlyFollowings[fid] = u
	}

	// get recent tweets
	for _, onlyFollow := range onlyFollowings {
		p := &sttypes.ListRecentInput{
			MaxResults:  10,
			Query:       "from:" + onlyFollow.Username + " -is:retweet -is:reply",
			TweetFields: fields.TweetFieldList{fields.TweetFieldCreatedAt},
		}
		res, err := searchtweet.ListRecent(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("----- %s's recent Tweets -----\n", onlyFollow.displayName())
		c := 0
		for _, t := range res.Data {
			if c > 3 {
				break
			}
			fmt.Printf("[%s] %s\n", t.CreatedAt, StringValue(t.Text))
			c++
		}

		fmt.Println()
	}
}
