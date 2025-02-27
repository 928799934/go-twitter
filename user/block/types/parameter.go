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

type ListInput struct {
	// Path parameter
	ID string

	// Query parameters
	MaxResults      ListMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (m ListMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m ListMaxResults) String() string {
	return strconv.Itoa(int(m))
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

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type CreateInput struct {
	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TargetID string `json:"target_user_id"` // required
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
	// Path parameters
	SourceUserID string // The authenticated user ID
	TargetID     string // The user ID for unfollow
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

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
