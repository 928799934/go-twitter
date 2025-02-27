package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/928799934/go-twitter/internal/util"
	jsoniter "github.com/json-iterator/go"
)

type ComplianceType string

const (
	ComplianceTypeTweets ComplianceType = "tweets"
	ComplianceTypeUsers  ComplianceType = "users"
)

type ComplianceStatus string

const (
	ComplianceStatusCreated    ComplianceStatus = "created"
	ComplianceStatusInProgress ComplianceStatus = "in_progress"
	ComplianceStatusFailed     ComplianceStatus = "failed"
	ComplianceStatusComplete   ComplianceStatus = "complete"
)

type ListJobsInput struct {
	// Query Parameters
	Type   ComplianceType
	Status ComplianceStatus
}

var listJobsQueryParameters = map[string]struct{}{
	"type":   {},
	"status": {},
}

func (p *ListJobsInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Type == "" {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listJobsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListJobsInput) Body() io.Reader {
	return nil
}

func (p *ListJobsInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *ListJobsInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["type"] = string(p.Type)

	if p.Status != "" {
		m["status"] = string(p.Status)
	}

	return m
}

type GetJobInput struct {
	// Path parameters
	ID string
}

func (p *GetJobInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *GetJobInput) Body() io.Reader {
	return nil
}

func (p *GetJobInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *GetJobInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type CreateJobInput struct {
	// JSON body parameter
	Type      ComplianceType `json:"type,omitempty"` // required
	Name      *string        `json:"name,omitempty"`
	Resumable *bool          `json:"resumable,omitempty"`
}

func (p *CreateJobInput) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *CreateJobInput) Body() io.Reader {
	jsonData, _ := jsoniter.MarshalToString(p)
	return strings.NewReader(jsonData)
}

func (p *CreateJobInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *CreateJobInput) ParameterMap() map[string]string {
	return map[string]string{}
}
