package pinnedlist

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/list/pinnedlist/types"
)

const (
	listEndpoint   = gotwitter.Domain + "/2/users/:id/pinned_lists"
	createEndpoint = gotwitter.Domain + "/2/users/:id/pinned_lists"
	deleteEndpoint = gotwitter.Domain + "/2/users/:id/pinned_lists/:list_id"
)

// Returns the Lists pinned by a specified user.
// https://developer.twitter.com/en/docs/twitter-api/lists/pinned-lists/api-reference/get-users-id-pinned_lists
func List(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to pin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-pinned-lists
func Create(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unpin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-pinned-lists-list_id
func Delete(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, http.MethodDelete, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
