package upload

import (
	"context"
	"errors"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/media/upload/types"
)

const (
	initializeEndpoint = gotwitter.Domain + "/2/media/upload/initialize"
	appendEndpoint     = gotwitter.Domain + "/2/media/upload/:mediaID/append"
	finalizeEndpoint   = gotwitter.Domain + "/2/media/upload/:mediaID/finalize"
)

func Initialize(ctx context.Context, c *gotwitter.GoTwitter, p *types.InitializeInput) (*types.InitializeOutput, error) {
	if p == nil {
		return nil, errors.New("InitializeInput is nil")
	}
	res := &types.InitializeOutput{}
	if err := c.CallAPI(ctx, initializeEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Append(ctx context.Context, c *gotwitter.GoTwitter, p *types.AppendInput) (*types.AppendOutput, error) {
	if p == nil {
		return nil, errors.New("AppendInput is nil")
	}
	res := &types.AppendOutput{}
	if err := c.CallAPI(ctx, appendEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Finalize(ctx context.Context, c *gotwitter.GoTwitter, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	if p == nil {
		return nil, errors.New("FinalizeInput is nil")
	}
	res := &types.FinalizeOutput{}
	if err := c.CallAPI(ctx, finalizeEndpoint, http.MethodPost, gotwitter.OAuth1UserContext, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
