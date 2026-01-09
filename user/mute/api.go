package mute

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/user/mute/types"
)

const (
	listEndpoint   = gotwitter.Domain + "/2/users/:id/muting"
	createEndpoint = gotwitter.Domain + "/2/users/:id/muting"
	deleteEndpoint = gotwitter.Domain + "/2/users/:source_user_id/muting/:target_user_id"
)

// Returns a list of users who are muted by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func Lists(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListsInput) (*types.ListsOutput, error) {
	res := &types.ListsOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to mute the target user.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
func Create(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to unmute the target user.
// The request succeeds with no action when the user sends a request to a user they're not muting or have already unmuted.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
func Delete(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, http.MethodDelete, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
