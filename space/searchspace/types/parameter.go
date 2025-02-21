package types

import (
	"io"
	"strconv"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
)

type ListMaxResults int

type ListInput struct {
	// Path parameters
	Query       string
	Expansions  fields.ExpansionList
	MaxResults  ListMaxResults
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
	State       fields.State
}

var listQueryParameters = map[string]struct{}{
	"query":        {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
	"max_results":  {},
	"state":        {},
}

func (m ListMaxResults) Valid() bool {
	return m > 0 && m <= 100
}

func (m ListMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

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

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)

	if p.State.Valid() {
		m["state"] = p.State.String()
	}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	return m
}
