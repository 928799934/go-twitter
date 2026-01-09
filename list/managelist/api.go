package managelist

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/list/managelist/types"
)

const (
	createEndpoint = gotwitter.Domain + "/2/lists"
	updateEndpoint = gotwitter.Domain + "/2/lists/:id"
	deleteEndpoint = gotwitter.Domain + "/2/lists/:id"
)

// Enables the authenticated user to create a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists
func Create(ctx context.Context, c *gotwitter.GoTwitter, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to update the meta data of a specified List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func Update(ctx context.Context, c *gotwitter.GoTwitter, p *types.UpdateInput) (*types.UpdateOutput, error) {
	res := &types.UpdateOutput{}
	if err := c.CallAPI(ctx, updateEndpoint, http.MethodPut, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to delete a List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func Delete(ctx context.Context, c *gotwitter.GoTwitter, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, http.MethodDelete, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
