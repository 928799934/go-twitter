package main

import (
	"context"

	gotwitter "github.com/928799934/go-twitter"
	. "github.com/928799934/go-twitter/_examples"
	"github.com/928799934/go-twitter/tweet/managetweet"
	"github.com/928799934/go-twitter/tweet/managetweet/types"
)

// SimpleTweet posts a tweet with only text, and return posted tweet ID.
func SimpleTweet(c *gotwitter.GoTwitter, text string) (string, error) {
	p := &types.CreateInput{
		Text: String(text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return StringValue(res.Data.ID), nil
}

// DeleteTweet deletes a tweet specified by tweet ID.
func DeleteTweet(c *gotwitter.GoTwitter, id string) (bool, error) {
	p := &types.DeleteInput{
		ID: id,
	}

	res, err := managetweet.Delete(context.Background(), c, p)
	if err != nil {
		return false, err
	}

	return BoolValue(res.Data.Deleted), nil
}
