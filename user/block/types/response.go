package types

import "github.com/928799934/go-twitter/resources"

type ListOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type CreateOutput struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
