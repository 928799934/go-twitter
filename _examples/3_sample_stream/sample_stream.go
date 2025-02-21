package main

import (
	"context"
	"fmt"

	gotwitter "github.com/928799934/go-twitter"
	. "github.com/928799934/go-twitter/_examples"
	"github.com/928799934/go-twitter/tweet/samplestream"
	"github.com/928799934/go-twitter/tweet/samplestream/types"
)

func samplingTweets(c *gotwitter.GoTwitter, count int) {
	p := &types.SampleStreamInput{}
	s, err := samplestream.SampleStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		cnt++
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(StringValue(t.Data.ID), StringValue(t.Data.Text))
		}

		if cnt > count {
			s.Stop()
			break
		}
	}
}
