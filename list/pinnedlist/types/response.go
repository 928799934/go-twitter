package types

import "github.com/928799934/go-twitter/resources"

type ListOutput struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
}

func (r *ListOutput) HasPartialError() bool {
	return false
}

type CreateOutput struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
