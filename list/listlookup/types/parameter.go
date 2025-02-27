package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
)

type GetInput struct {
	// Path parameter
	ID string

	// Query parameters
	Expansions fields.ExpansionList
	ListFields fields.ListFieldList
	UserFields fields.UserFieldList
}

var getQueryParameters = map[string]struct{}{
	"expansions":  {},
	"list.fields": {},
	"user.fields": {},
}

func (p *GetInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetInput) Body() io.Reader {
	return nil
}

func (p *GetInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *GetInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)
	return m
}

type ListOwnedMaxResults int

func (m ListOwnedMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListOwnedMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListOwnedInput struct {
	// Path parameter
	ID string // User ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListOwnedMaxResults
	PaginationToken string
}

var listOwnedQueryParameters = map[string]struct{}{
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *ListOwnedInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listOwnedQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListOwnedInput) Body() io.Reader {
	return nil
}

func (p *ListOwnedInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *ListOwnedInput) ParameterMap() map[string]string {
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
