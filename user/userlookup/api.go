package userlookup

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/user/userlookup/types"
)

const (
	listEndpoint            = "https://api.twitter.com/2/users"
	getEndpoint             = "https://api.twitter.com/2/users/:id"
	listByUsernamesEndpoint = "https://api.twitter.com/2/users/by"
	getByUsernameEndpoint   = "https://api.twitter.com/2/users/by/username/:username"
	getMeEndpoint           = "https://api.twitter.com/2/users/me"
)

// GET /2/users
// Returns a variety of information about one or more users specified by the requested IDs.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func List(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/:id
// Returns a variety of information about a single user specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func Get(ctx context.Context, c *gotwitter.GoTwitter, p *types.GetInput) (*types.GetOutput, error) {
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/by
// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func ListByUsernames(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListByUsernamesInput) (*types.ListByUsernamesOutput, error) {
	res := &types.ListByUsernamesOutput{}
	if err := c.CallAPI(ctx, listByUsernamesEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/by/username/:username
// Returns a variety of information about a single user specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func GetByUsername(ctx context.Context, c *gotwitter.GoTwitter, p *types.GetByUsernameInput) (*types.GetByUsernameOutput, error) {
	res := &types.GetByUsernameOutput{}
	if err := c.CallAPI(ctx, getByUsernameEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/me
// Returns information about an authorized user.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func GetMe(ctx context.Context, c *gotwitter.GoTwitter, p *types.GetMeInput) (*types.GetMeOutput, error) {
	res := &types.GetMeOutput{}
	if err := c.CallAPI(ctx, getMeEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
