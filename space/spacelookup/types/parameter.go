package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
)

type GetInput struct {
	// Path parameter
	ID string // required: Space ID

	// Query parameters
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var getQueryParameters = map[string]struct{}{
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
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

func (p *GetInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}

// ListInput is struct of parameters
// for request GET /2/spaces
type ListInput struct {
	// Query parameters
	IDs         []string // required: Space IDs
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var listQueryParameters = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	if len(p.IDs) == 0 {
		return ""
	}

	endpoint := endpointBase

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
	m["ids"] = util.QueryValue(p.IDs)
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}

// ListByCreatorIDsInput is struct of parameters
// for request GET /2/spaces/by/creator_ids
type ListByCreatorIDsInput struct {
	// Query parameters
	UserIDs     []string // required
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var listByCreatorIDsQueryParameters = map[string]struct{}{
	"user_ids":     {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *ListByCreatorIDsInput) ResolveEndpoint(endpointBase string) string {
	if len(p.UserIDs) == 0 {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listByCreatorIDsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListByCreatorIDsInput) Body() io.Reader {
	return nil
}

func (p *ListByCreatorIDsInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m["user_ids"] = util.QueryValue(p.UserIDs)
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}

type ListBuyersInput struct {
	// Path parameter
	ID string // required: Space ID

	// Query parameters
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var listBuyersQueryParameters = map[string]struct{}{
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *ListBuyersInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listBuyersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListBuyersInput) Body() io.Reader {
	return nil
}

func (p *ListBuyersInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)
	return m
}

type ListTweetsInput struct {
	// Path parameter
	ID string

	// Query parameters
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var listTweetsQueryParameters = map[string]struct{}{
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *ListTweetsInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listTweetsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListTweetsInput) Body() io.Reader {
	return nil
}

func (p *ListTweetsInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)
	return m
}
