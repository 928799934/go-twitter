package types

import (
	"io"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type UpdateInput struct {
	// Path parameter
	ID string `json:"-"` // The tweet ID to hide or unhide

	// JSON body parameter
	Hidden bool `json:"hidden"` // required
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

func (p *UpdateInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *UpdateInput) ParameterMap() map[string]string {
	return map[string]string{}
}
