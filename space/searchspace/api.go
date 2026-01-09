package searchspace

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/space/searchspace/types"
)

const (
	listEndpoint = gotwitter.Domain + "/2/spaces/search"
)

// Return live or scheduled Spaces matching your specified search terms.
// This endpoint performs a keyword search, meaning that it will return Spaces
// that are an exact case-insensitive match of the specified search term. The search term will match the original title of the Space.
// https://developer.twitter.com/en/docs/twitter-api/spaces/search/api-reference/get-spaces-search
func List(ctx context.Context, c *gotwitter.GoTwitter, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
