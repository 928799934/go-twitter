package media

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/media/types"
)

const (
	initEndpoint     = gotwitter.Domain + "/2/media/upload"
	appendEndpoint   = gotwitter.Domain + "/2/media/upload"
	finalizeEndpoint = gotwitter.Domain + "/2/media/upload"
	statusEndpoint   = gotwitter.Domain + "/2/media/upload"
)

func Init(ctx context.Context, c *gotwitter.GoTwitter, p *types.InitInput) (*types.InitOutput, error) {
	res := &types.InitOutput{}
	if err := c.CallAPI(ctx, initEndpoint, http.MethodPost, gotwitter.OAuth2AccessToken, p, res); err != nil {
		return nil, err
	}
	return res, nil
}

func Append(ctx context.Context, c *gotwitter.GoTwitter, p *types.AppendInput) error {
	if err := c.CallAPI(ctx, appendEndpoint, http.MethodPost, gotwitter.OAuth2AccessToken, p, nil); err != nil {
		return err
	}
	return nil
}

func Finalize(ctx context.Context, c *gotwitter.GoTwitter, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	res := &types.FinalizeOutput{}
	if err := c.CallAPI(ctx, finalizeEndpoint, http.MethodPost, gotwitter.OAuth2AccessToken, p, res); err != nil {
		return nil, err
	}
	return res, nil
}

func Status(ctx context.Context, c *gotwitter.GoTwitter, p *types.StatusInput) (*types.StatusOutput, error) {
	res := &types.StatusOutput{}
	if err := c.CallAPI(ctx, statusEndpoint, http.MethodGet, gotwitter.OAuth2AccessToken, p, res); err != nil {
		return nil, err
	}
	return res, nil
}
