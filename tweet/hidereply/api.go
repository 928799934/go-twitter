package hidereply

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/hidereply/types"
)

const updateEndpoint = "https://api.twitter.com/2/tweets/:id/hidden"

// Hides or unhides a reply to a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func Update(ctx context.Context, c *gotwitter.GoTwitter, p *types.UpdateInput) (*types.UpdateOutput, error) {
	res := &types.UpdateOutput{}
	if err := c.CallAPI(ctx, updateEndpoint, http.MethodPut, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
