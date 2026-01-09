package types

import (
	"bytes"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type MediaCategory string

const (
	MediaCategoryAmplifyVideo MediaCategory = "amplify_video"
	MediaCategoryTweetGIF     MediaCategory = "tweet_gif"
	MediaCategoryTweetImage   MediaCategory = "tweet_image"
	MediaCategoryTweetVideo   MediaCategory = "tweet_video"
	MediaCategoryDMGIF        MediaCategory = "dm_gif"
	MediaCategoryDMImage      MediaCategory = "dm_image"
	MediaCategoryDMVideo      MediaCategory = "dm_video"
	MediaCategorySubtitles    MediaCategory = "subtitles"
)

type MediaType string

const (
	MediaTypeMP4       MediaType = "video/mp4"
	MediaTypeWebM      MediaType = "video/webm"
	MediaTypeMP2T      MediaType = "video/mp2t"
	MediaTypeQuickTime MediaType = "video/quicktime"
	MediaTypeSRT       MediaType = "text/srt"
	MediaTypeVTT       MediaType = "text/vtt"
	MediaTypeJPEG      MediaType = "image/jpeg"
	MediaTypeGIF       MediaType = "image/gif"
	MediaTypeBMP       MediaType = "image/bmp"
	MediaTypePNG       MediaType = "image/png"
	MediaTypeWebP      MediaType = "image/webp"
	MediaTypePJPEG     MediaType = "image/pjpeg"
	MediaTypeTIFF      MediaType = "image/tiff"
	MediaTypeGLTF      MediaType = "model/gltf-binary"
	MediaTypeUSDZ      MediaType = "model/vnd.usdz+zip"
)

// type Parameters interface {
// 	ResolveEndpoint(endpointBase string) string
// 	Body() io.Reader
// 	ContentType() string
// 	ParameterMap() map[string]string
// }

// InitializeInput is the input for the Initialize endpoint.
type InitializeInput struct {
	// Unique identifier of this User. This is returned as a string in order to avoid complications
	// with languages and tools that cannot handle large integers.
	AdditionalOwners []string `json:"additional_owners,omitempty"`

	// A string enum value which identifies a media use-case. This identifier is used to enforce use-case specific constraints
	// (e.g. file size, video duration) and enable advanced features.
	MediaCategory MediaCategory `json:"media_category,omitempty"`

	//The type of media.
	MediaType MediaType `json:"media_type,omitempty"`

	// Whether this media is shared or not.
	Shared bool `json:"shared,omitempty"`

	// The total size of the media upload in bytes.
	TotalBytes int `json:"total_bytes,omitempty"`
}

func (p *InitializeInput) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *InitializeInput) Body() io.Reader {
	jsonData, err := jsoniter.MarshalToString(p)
	if err != nil {
		return nil
	}
	return strings.NewReader(jsonData)
}

func (p *InitializeInput) ParameterMap() map[string]string {
	return map[string]string{}
}

func (p *InitializeInput) ContentType() string {
	return "application/json;charset=UTF-8"
}

type AppendInput struct {
	// Path parameter: The media identifier for the media to perform the append operation.
	MediaID string

	// The file to upload.
	Media io.Reader

	// An integer value representing the media upload segment.
	SegmentIndex int

	w *multipart.Writer
}

func (p *AppendInput) ResolveEndpoint(endpointBase string) string {
	endpoint := strings.Replace(endpointBase, ":mediaID", p.MediaID, 1)
	return endpoint
}

func (p *AppendInput) Body() io.Reader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(p.w.Boundary())
	part, _ := writer.CreateFormField("media")
	io.Copy(part, p.Media)
	writer.WriteField("segment_index", strconv.Itoa(p.SegmentIndex))
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
	return map[string]string{}
}

type FinalizeInput struct {

	// Path parameter: The media identifier for the media to perform the finalize operation.
	MediaID string
}

func (p *FinalizeInput) ResolveEndpoint(endpointBase string) string {
	endpoint := strings.Replace(endpointBase, ":mediaID", p.MediaID, 1)
	return endpoint
}

func (p *FinalizeInput) Body() io.Reader {
	return nil
}

func (p *FinalizeInput) ParameterMap() map[string]string {
	return map[string]string{}
}

func (p *FinalizeInput) ContentType() string {
	return "application/json;charset=UTF-8"
}
