package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
	jsoniter "github.com/json-iterator/go"
)

type ListInput struct {
	// Path parameter
	ID string // User ID

	// Query parameter
	Expansions fields.ExpansionList
	ListFields fields.ListFieldList
	UserFields fields.UserFieldList
}

var listQueryParameters = map[string]struct{}{
	"expansions":  {},
	"list.fields": {},
	"user.fields": {},
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListInput) Body() io.Reader {
	return nil
}

func (p *ListInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)
	return m
}

type CreateInput struct {
	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID string `json:"list_id"` // required
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *CreateInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *CreateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteInput struct {
	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *DeleteInput) Body() io.Reader {
	return nil
}

func (p *DeleteInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
