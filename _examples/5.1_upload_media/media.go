package main

import (
	"context"

	gotwitter "github.com/928799934/go-twitter"
	. "github.com/928799934/go-twitter/_examples"
	"github.com/928799934/go-twitter/media/upload"
	"github.com/928799934/go-twitter/media/upload/types"
	"github.com/928799934/go-twitter/tweet/managetweet"
	mtTypes "github.com/928799934/go-twitter/tweet/managetweet/types"
)

func Initialize(c *gotwitter.GoTwitter, p *types.InitializeInput) (*types.InitializeOutput, error) {
	res, err := upload.Initialize(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Append(c *gotwitter.GoTwitter, p *types.AppendInput) (*types.AppendOutput, error) {
	res, err := upload.Append(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Finalize(c *gotwitter.GoTwitter, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	res, err := upload.Finalize(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func PostWithMedia(c *gotwitter.GoTwitter, text string, mediaID string) (string, error) {
	p := &mtTypes.CreateInput{
		Text: String(text),
		Media: &mtTypes.CreateInputMedia{
			MediaIDs: []string{mediaID},
		},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return StringValue(res.Data.ID), nil
}
