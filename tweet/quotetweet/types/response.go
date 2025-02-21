package types

import "github.com/928799934/go-twitter/resources"

type ListOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users,omitempty"`
	} `json:"includes,omitempty"`
	Meta   resources.QuoteTweetsMeta `json:"meta"`
	Errors []resources.PartialError  `json:"errors,omitempty"`
}

func (r *ListOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
