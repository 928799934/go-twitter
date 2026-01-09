package samplestream

import (
	"context"
	"net/http"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/samplestream/types"
)

const (
	sampleStreamEndpoint = gotwitter.Domain + "/2/tweets/sample/stream"
)

// Streams about 1% of all Tweets in real-time.
// If you have Academic Research access, you can connect up to two redundant connections to maximize your streaming up-time.
// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/api-reference/get-tweets-sample-stream
func SampleStream(ctx context.Context, c *gotwitter.GoTwitter, p *types.SampleStreamInput) (*gotwitter.StreamClient[*types.SampleStreamOutput], error) {
	tc := gotwitter.NewTypedClient[*types.SampleStreamOutput](c)
	s, err := tc.CallStream(ctx, sampleStreamEndpoint, http.MethodGet, gotwitter.OAuth2BearerToken, p)
	if err != nil {
		return nil, err
	}

	return s, nil
}
