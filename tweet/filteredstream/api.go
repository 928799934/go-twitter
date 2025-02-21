package filteredstream

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/filteredstream/types"
)

const (
	listRulesEndpoint           = "https://api.twitter.com/2/tweets/search/stream/rules"
	createOrDeleteRulesEndpoint = "https://api.twitter.com/2/tweets/search/stream/rules"
	searchStreamEndpoint        = "https://api.twitter.com/2/tweets/search/stream"
)

// Return a list of rules currently active on the streaming endpoint, either as a list or individually.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
func ListRules(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListRulesInput) (*types.ListRulesOutput, error) {
	res := &types.ListRulesOutput{}
	if err := c.CallAPI(ctx, listRulesEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Add rules to your stream.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules
func CreateRules(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateRulesInput) (*types.CreateRulesOutput, error) {
	res := &types.CreateRulesOutput{}
	if err := c.CallAPI(ctx, createOrDeleteRulesEndpoint, http.MethodPost, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Delete rules to your stream.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/post-tweets-search-stream-rules
func DeleteRules(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteRulesInput) (*types.DeleteRulesOutput, error) {
	res := &types.DeleteRulesOutput{}
	if err := c.CallAPI(ctx, createOrDeleteRulesEndpoint, http.MethodPost, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Streams Tweets in real-time that match the rules that you added to the stream using the POST /tweets/search/stream/rules endpoint.
// If you haven't added any rules to your stream, you will not receive any Tweets.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream
func SearchStream(ctx context.Context, c *gotwitter.GoTwitter, p *types.SearchStreamInput) (*gotwitter.StreamClient[*types.SearchStreamOutput], error) {
	tc := gotwitter.NewTypedClient[*types.SearchStreamOutput](c)
	s, err := tc.CallStream(ctx, searchStreamEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p)
	if err != nil {
		return nil, err
	}

	return s, nil
}
