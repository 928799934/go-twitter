package types

import "github.com/928799934/go-twitter/resources"

type GetOutput struct {
	Data     resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *GetOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type ListOwnedOutput struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes,omitempty"`
	Meta   resources.ListLookupOwnedListsMeta
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *ListOwnedOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
