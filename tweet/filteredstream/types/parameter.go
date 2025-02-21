package types

import (
	"io"
	"strconv"
	"strings"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
	jsoniter "github.com/json-iterator/go"
)

type ListRulesInput struct {
	// Query parameters
	IDs []string
}

var listRulesQueryParameters = map[string]struct{}{
	"ids": {},
}

func (p *ListRulesInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListRulesInput) Body() io.Reader {
	return nil
}

func (p *ListRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if len(p.IDs) > 0 {
		m["ids"] = util.QueryValue(p.IDs)
	}

	return m
}

type AddingRules []AddingRule

type AddingRule struct {
	Value *string `json:"value,omitempty"`
	Tag   *string `json:"tag,omitempty"`
}

type DeletingRules struct {
	IDs []string `json:"ids"`
}

type CreateRulesInput struct {
	// Query parameters
	DryRun bool `json:"-"` // default false

	// JSON body parameter
	Add AddingRules `json:"add,omitempty"`
}

var createOrDeleteRulesQueryParameters = map[string]struct{}{
	"dry_run": {},
}

func (p *CreateRulesInput) ResolveEndpoint(endpointBase string) string {
	if len(p.Add) == 0 {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, createOrDeleteRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *CreateRulesInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *CreateRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m["dry_run"] = strconv.FormatBool(p.DryRun)
	return m
}

type DeleteRulesInput struct {
	// Query parameters
	DryRun bool `json:"-"` // default false

	// JSON body parameter
	Delete *DeletingRules `json:"delete,omitempty"`
}

func (p *DeleteRulesInput) ResolveEndpoint(endpointBase string) string {
	if p.Delete == nil {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, createOrDeleteRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *DeleteRulesInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *DeleteRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m["dry_run"] = strconv.FormatBool(p.DryRun)
	return m
}

type SearchStreamBackfillMinutes int

func (s SearchStreamBackfillMinutes) Valid() bool {
	return int(s) > 0
}

func (s SearchStreamBackfillMinutes) String() string {
	return strconv.Itoa(int(s))
}

type SearchStreamInput struct {
	// Query parameters
	BackfillMinutes SearchStreamBackfillMinutes
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var searchStreamQueryParameters = map[string]struct{}{
	"backfill_minutes": {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *SearchStreamInput) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, searchStreamQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *SearchStreamInput) Body() io.Reader {
	return nil
}

func (p *SearchStreamInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.BackfillMinutes.Valid() {
		m["backfill_minutes"] = p.BackfillMinutes.String()
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}
