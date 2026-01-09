package block

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/user/block/types"
)

const (
	listEndpoint   = gotwitter.Domain + "/2/users/:id/blocking"
	createEndpoint = gotwitter.Domain + "/2/users/:id/blocking"
	deleteEndpoint = gotwitter.Domain + "/2/users/:source_user_id/blocking/:target_user_id"
)

// Returns a list of users who are blocked by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/get-users-blocking
func List(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user (in the path) to block the target user. The user (in the path) must match the user context authorizing the request.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/post-users-user_id-blocking
func Create(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to unblock another user.
// The request succeeds with no action when the user sends a request to a user they're not blocking or have already unblocked.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/delete-users-user_id-blocking
func Delete(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, http.MethodDelete, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
