package tweetlookup

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/tweetlookup/types"
)

const (
	listEndpoint = gotwitter.Domain + "/2/tweets"
	getEndpoint  = gotwitter.Domain + "/2/tweets/:id"
)

// Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func List(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single Tweet specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func Get(ctx context.Context, c *gotwitter.GoTwitter, p *types.GetInput) (*types.GetOutput, error) {
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
