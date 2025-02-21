package types

import "github.com/928799934/go-twitter/resources"

type ListOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users"`
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListTweetsLookupMeta `json:"meta"`
	Errors []resources.PartialError       `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
