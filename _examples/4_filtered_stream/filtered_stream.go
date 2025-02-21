package main

import (
	"context"
	"fmt"

	. "github.com/928799934/go-twitter/_examples"
	"github.com/928799934/go-twitter/tweet/filteredstream"
	"github.com/928799934/go-twitter/tweet/filteredstream/types"
)

// createSearchStreamRules lists search stream rules.
func listSearchStreamRules() {
	c := newOAuth2Client()

	p := &types.ListRulesInput{}
	res, err := filteredstream.ListRules(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", StringValue(r.ID), StringValue(r.Value), StringValue(r.Tag))
	}
}

func deleteSearchStreamRules(ruleID string) {
	c := newOAuth2Client()

	p := &types.DeleteRulesInput{
		Delete: &types.DeletingRules{
			IDs: []string{
				ruleID,
			},
		},
	}

	res, err := filteredstream.DeleteRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", StringValue(r.ID), StringValue(r.Value), StringValue(r.Tag))
	}
}

// createSearchStreamRules creates a search stream rule.
func createSearchStreamRules(keyword string) {
	c := newOAuth2Client()

	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: String(keyword), Tag: String(keyword)},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", StringValue(r.ID), StringValue(r.Value), StringValue(r.Tag))
	}
}

// execSearchStream call GET /2/tweets/search/stream API
// and outputs up to 10 results.
func execSearchStream() {
	c := newOAuth2Client()

	p := &types.SearchStreamInput{}
	s, err := filteredstream.SearchStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			if t != nil {
				cnt++
				fmt.Println(StringValue(t.Data.ID), StringValue(t.Data.Text))
			}
		}

		if cnt > 10 {
			s.Stop()
			break
		}
	}
}
