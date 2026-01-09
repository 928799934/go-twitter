package listfollow

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/list/listfollow/types"
)

const (
	listFollowersEndpoint = gotwitter.Domain + "/2/lists/:id/followers"
	listFollowedEndpoint  = gotwitter.Domain + "/2/users/:id/followed_lists"
	createEndpoint        = gotwitter.Domain + "/2/users/:id/followed_lists"
	deleteEndpoint        = gotwitter.Domain + "/2/users/:id/followed_lists/:list_id"
)

// Returns a list of users who are followers of the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-lists-id-followers
func ListFollowers(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListFollowersInput) (*types.ListFollowersOutput, error) {
	res := &types.ListFollowersOutput{}
	if err := c.CallAPI(ctx, listFollowersEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists a specified user follows.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-users-id-followed_lists
func ListFollowed(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListFollowedInput) (*types.ListFollowedOutput, error) {
	res := &types.ListFollowedOutput{}
	if err := c.CallAPI(ctx, listFollowedEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to follow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-followed-lists
func Create(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unfollow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func Delete(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, http.MethodDelete, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
