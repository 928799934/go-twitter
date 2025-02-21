package types

import "github.com/928799934/go-twitter/resources"

type ListRecentOutput struct {
	Data   []resources.TweetCount         `json:"data"`
	Meta   resources.TweetCountRecentMeta `json:"meta"`
	Errors []resources.PartialError       `json:"errors"`
}

func (r *ListRecentOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type ListAllOutput struct {
	Data   []resources.TweetCount      `json:"data"`
	Meta   resources.TweetCountAllMeta `json:"meta"`
	Errors []resources.PartialError    `json:"errors"`
}

func (r *ListAllOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
