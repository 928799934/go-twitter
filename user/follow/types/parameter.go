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

type ListMaxResults int

func (m ListMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m ListMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListFollowingsInput struct {
	// Path parameter
	ID string

	// Query parameters
	MaxResults      ListMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listFollowingsQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListFollowingsInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listFollowingsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListFollowingsInput) Body() io.Reader {
	return nil
}

func (p *ListFollowingsInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	if len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions.Values())
	}

	if len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields.Values())
	}

	if len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields.Values())
	}

	return m
}

type ListFollowersInput struct {
	// Path parameter
	ID string

	// Query parameters
	MaxResults      ListMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listFollowersQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListFollowersInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listFollowersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListFollowersInput) Body() io.Reader {
	return nil
}

func (p *ListFollowersInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	if len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions.Values())
	}

	if len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields.Values())
	}

	if len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields.Values())
	}

	return m
}

type CreateFollowingInput struct {
	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TargetID string `json:"target_user_id"` // required
}

func (p *CreateFollowingInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateFollowingInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *CreateFollowingInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteFollowingInput struct {
	// Path parameters
	SourceUserID string // The authenticated user ID
	TargetID     string // The user ID for unfollow
}

func (p *DeleteFollowingInput) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

	return endpoint
}

func (p *DeleteFollowingInput) Body() io.Reader {
	return nil
}

func (p *DeleteFollowingInput) ParameterMap() map[string]string {
	return map[string]string{}
}
