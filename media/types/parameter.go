package types

import (
	"bytes"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/internal/util"
)

type InitInput struct {
	// Query parameters
	TotalBytes    int                  // required
	MediaType     string               // required
	MediaCategory fields.MediaCategory // required
}

var initQueryParameters = map[string]struct{}{
	"media_category": {},
	"total_bytes":    {},
	"media_type":     {},
	"command":        {},
}

func (p *InitInput) ResolveEndpoint(endpointBase string) string {

	if p.MediaType == "" || p.TotalBytes == 0 || p.MediaCategory == "" {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, initQueryParameters)
		endpoint += "?" + qs
	}
	return endpoint
}

func (p *InitInput) Body() io.Reader {
	return nil
}

func (p *InitInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *InitInput) ParameterMap() map[string]string {
	m := map[string]string{
		"command": "INIT",
	}
	if p.MediaType != "" {
		m["media_type"] = p.MediaType
	}
	if p.TotalBytes != 0 {
		m["total_bytes"] = strconv.Itoa(p.TotalBytes)
	}

	if p.MediaCategory != "" {
		m["media_category"] = p.MediaCategory.String()
	}
	return m
}

type AppendInput struct {
	// Query parameters
	MediaID      string // required
	Data         []byte // required
	SegmentIndex int

	w *multipart.Writer
}

var appendQueryParameters = map[string]struct{}{
	"media_id":      {},
	"segment_index": {},
	"command":       {},
}

func (p *AppendInput) ResolveEndpoint(endpointBase string) string {

	if p.MediaID == "" || p.Data == nil {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, appendQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *AppendInput) Body() io.Reader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(p.w.Boundary())
	part, _ := writer.CreateFormField("media")
	io.Copy(part, bytes.NewReader(p.Data))
	writer.Close()
	return body
}

func (p *AppendInput) ContentType() string {
	if p.w == nil {
		p.w = multipart.NewWriter(nil)
	}
	return p.w.FormDataContentType()
}

func (p *AppendInput) ParameterMap() map[string]string {
	m := map[string]string{
		"command": "APPEND",
	}

	if p.MediaID != "" {
		m["media_id"] = p.MediaID
	}

	m["segment_index"] = strconv.Itoa(p.SegmentIndex)
	return m
}

// FINALIZE
type FinalizeInput struct {
	// Query parameters
	MediaID string // required
}

var finalizeQueryParameters = map[string]struct{}{
	"media_id": {},
	"command":  {},
}

func (p *FinalizeInput) ResolveEndpoint(endpointBase string) string {

	if p.MediaID == "" {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, finalizeQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *FinalizeInput) Body() io.Reader {
	return nil
}

func (p *FinalizeInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *FinalizeInput) ParameterMap() map[string]string {
	m := map[string]string{
		"command": "FINALIZE",
	}
	if p.MediaID != "" {
		m["media_id"] = p.MediaID
	}
	return m
}

type StatusInput struct {
	// Query parameters
	MediaID string // required
}

var statusQueryParameters = map[string]struct{}{
	"media_id": {},
	"command":  {},
}

func (p *StatusInput) ResolveEndpoint(endpointBase string) string {

	if p.MediaID == "" {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, statusQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *StatusInput) Body() io.Reader {
	return nil
}

func (p *StatusInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

func (p *StatusInput) ParameterMap() map[string]string {
	m := map[string]string{
		"command": "STATUS",
	}
	if p.MediaID != "" {
		m["media_id"] = p.MediaID
	}
	return m
}
