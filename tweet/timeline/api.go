package timeline

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/timeline/types"
)

const (
	listTweetsEndpoint               = gotwitter.Domain + "/2/users/:id/tweets"
	listMentionsEndpoint             = gotwitter.Domain + "/2/users/:id/mentions"
	listReverseChronologicalEndpoint = gotwitter.Domain + "/2/users/:id/timelines/reverse_chronological"
)

// Returns Tweets composed by a single user, specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request. Using pagination, the most recent 3,200 Tweets can be retrieved.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func ListTweets(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListTweetsInput) (*types.ListTweetsOutput, error) {
	res := &types.ListTweetsOutput{}
	if err := c.CallAPI(ctx, listTweetsEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns Tweets mentioning a single user specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request. Using pagination, up to the most recent 800 Tweets can be retrieved.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions
func ListMentions(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListMentionsInput) (*types.ListMentionsOutput, error) {
	res := &types.ListMentionsOutput{}
	if err := c.CallAPI(ctx, listMentionsEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows you to retrieve a collection of the most recent Tweets and Retweets
// posted by you and users you follow. This endpoint can return every Tweet
// created on a timeline over the last 7 days as well as the most recent 800 regardless of creation date.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-reverse-chronological
func ListReverseChronological(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListReverseChronologicalInput) (*types.ListReverseChronologicalOutput, error) {
	res := &types.ListReverseChronologicalOutput{}
	if err := c.CallAPI(ctx, listReverseChronologicalEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
