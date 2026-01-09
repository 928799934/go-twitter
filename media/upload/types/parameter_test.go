package types

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InitializeInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *InitializeInput
		expect io.Reader
	}{
		{
			name: "ok: has all parameters",
			params: &InitializeInput{
				AdditionalOwners: []string{"owner1", "owner2"},
				MediaCategory:    MediaCategoryTweetImage,
				MediaType:        MediaTypeJPEG,
				Shared:           true,
				TotalBytes:       1024,
			},
			expect: strings.NewReader(`{"additional_owners":["owner1","owner2"],"media_category":"tweet_image","media_type":"image/jpeg","shared":true,"total_bytes":1024}`),
		},
		{
			name: "ok: has some parameters",
			params: &InitializeInput{
				MediaCategory: MediaCategoryTweetImage,
				MediaType:     MediaTypeJPEG,
				TotalBytes:    1024,
			},
			expect: strings.NewReader(`{"media_category":"tweet_image","media_type":"image/jpeg","total_bytes":1024}`),
		},
		{
			name:   "ok: has no parameters",
			params: &InitializeInput{},
			expect: strings.NewReader(`{}`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Equal(tt, c.expect, r)
		})
	}
}

func Test_InitializeInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *InitializeInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &InitializeInput{MediaCategory: MediaCategoryTweetImage},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &InitializeInput{},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_InitializeInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload",
			expect:       "https://api.twitter.com/2/media/upload",
		},
		{
			name:         "empty",
			endpointBase: "",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &InitializeInput{}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
		})
	}
}

func Test_AppendInput_Body(t *testing.T) {
	cases := []struct {
		name     string
		params   *AppendInput
		boundary string
		wantErr  bool
		expect   string
	}{
		{
			name: "error: boundary is not set",
			params: &AppendInput{
				MediaID:      "test-media-id",
				Media:        strings.NewReader("test-media"),
				SegmentIndex: 1,
			},
			wantErr: true,
		},
		{
			name: "ok",
			params: &AppendInput{
				MediaID:      "test-media-id",
				Media:        strings.NewReader("test-media"),
				SegmentIndex: 1,
			},
			boundary: "test-boundary",
			wantErr:  false,
			expect:   "--test-boundary\r\nContent-Disposition: form-data; name=\"media\"; filename=\"media\"\r\nContent-Type: application/octet-stream\r\n\r\ntest-media\r\n--test-boundary\r\nContent-Disposition: form-data; name=\"segment_index\"\r\n\r\n1\r\n--test-boundary--\r\n",
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			b := c.params.Body()
			if c.wantErr {
				return
			}

			buf := new(bytes.Buffer)
			buf.ReadFrom(b)
			asst.Equal(c.expect, buf.String())
		})
	}
}

func Test_AppendInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *AppendInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &AppendInput{MediaID: "test-media-id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &AppendInput{},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_AppendInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		mediaID      string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "test-media-id",
			expect:       "https://api.twitter.com/2/media/upload/test-media-id",
		},
		{
			name:         "empty mediaID",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "",
			expect:       "https://api.twitter.com/2/media/upload/",
		},
		{
			name:         "empty endpoint",
			endpointBase: "",
			mediaID:      "test-media-id",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &AppendInput{
				MediaID: c.mediaID,
			}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
		})
	}
}

func Test_FinalizeInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *FinalizeInput
		expect io.Reader
	}{
		{
			name:   "normal: has parameters",
			params: &FinalizeInput{MediaID: "test-media-id"},
			expect: nil,
		},
		{
			name:   "normal: has no parameters",
			params: &FinalizeInput{},
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Equal(tt, c.expect, r)
		})
	}
}

func Test_FinalizeInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *FinalizeInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &FinalizeInput{MediaID: "test-media-id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &FinalizeInput{},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_FinalizeInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		mediaID      string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "test-media-id",
			expect:       "https://api.twitter.com/2/media/upload/test-media-id",
		},
		{
			name:         "empty mediaID",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "",
			expect:       "https://api.twitter.com/2/media/upload/",
		},
		{
			name:         "empty endpoint",
			endpointBase: "",
			mediaID:      "test-media-id",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &FinalizeInput{
				MediaID: c.mediaID,
			}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
		})
	}
}
