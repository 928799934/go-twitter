package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
	jsoniter "github.com/json-iterator/go"
)

type ListMembershipsMaxResults int

func (m ListMembershipsMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembershipsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListMembershipsInput struct {
	// Path parameter
	ID string // User ID

	// Query parameters
	MaxResults      ListMembershipsMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
}

var ListMembersListMembershipsQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
}

func (p *ListMembershipsInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, ListMembersListMembershipsQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListMembershipsInput) Body() io.Reader {
	return nil
}

func (p *ListMembershipsInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *ListMembershipsInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)

	return m
}

type ListMembersGetMaxResults int

func (m ListMembersGetMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembersGetMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListInput struct {
	// Path parameter
	ID string // List ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListMembersGetMaxResults
	PaginationToken string
}

var listQueryParameters = map[string]struct{}{
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

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

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	return m
}

type CreateInput struct {
	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	UserID string `json:"user_id,"` // required
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
	ID     string // List ID
	UserID string
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.UserID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedUserID := url.QueryEscape(p.UserID)
	endpoint = strings.Replace(endpoint, ":user_id", escapedUserID, 1)

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
