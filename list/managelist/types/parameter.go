package types

import (
	"io"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type CreateInput struct {
	// JSON body parameter
	Name        string  `json:"name,"` // required
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.Name == "" {
		return ""
	}

	return endpointBase
}

func (p *CreateInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *CreateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type UpdateInput struct {
	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *UpdateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *UpdateInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *UpdateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteInput struct {
	// Path parameter
	ID string // List ID
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *DeleteInput) Body() io.Reader {
	return nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
