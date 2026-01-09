package tweetcount

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/tweetcount/types"
)

const (
	listRecentEndpoint = gotwitter.Domain + "/2/tweets/counts/recent"
	listAllEndpoint    = gotwitter.Domain + "/2/tweets/counts/all"
)

// The recent Tweet counts endpoint returns count of Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func ListRecent(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListRecentInput) (*types.ListRecentOutput, error) {
	res := &types.ListRecentOutput{}
	if err := c.CallAPI(ctx, listRecentEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func ListAll(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListAllInput) (*types.ListAllOutput, error) {
	res := &types.ListAllOutput{}
	if err := c.CallAPI(ctx, listAllEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
