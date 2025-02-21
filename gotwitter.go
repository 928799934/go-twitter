package gotwitter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/928799934/go-twitter/internal/util"
	"github.com/928799934/go-twitter/resources"
	jsoniter "github.com/json-iterator/go"
)

type HTTPMethod = string

type AuthMethod string

const (
	OAuth1UserContext AuthMethod = "OAuth 1.0a User context"
	OAuth2BearerToken AuthMethod = "OAuth 2.0 Bearer token"
)

type GoTwitter struct {
	bearerToken string

	apiKey, apiKeySecret           string
	accessToken, accessTokenSecret string

	clientID, clientSecret string
}

func NewGoTwitter(opts ...OptionFunc) *GoTwitter {
	twitter := &GoTwitter{}

	for _, opt := range opts {
		opt(twitter)
	}

	return twitter
}

func (c *GoTwitter) CallAPI(ctx context.Context, uri string, method HTTPMethod, auth AuthMethod, p util.Parameters, i util.Response) error {

	var (
		header   map[string]string
		err      error
		jsonData []byte
		resp     *http.Response
	)

	uri = p.ResolveEndpoint(uri)

	switch auth {
	case OAuth1UserContext:
		params := url.Values{}
		for k, v := range p.ParameterMap() {
			params.Set(k, v)
		}
		if header, err = getOAuth1Header(c, method, uri, params); err != nil {
			return err
		}
	case OAuth2BearerToken:
		if header, err = getOAuth2Header(c); err != nil {
			return err
		}
	default:
		return fmt.Errorf("AuthMethod not support")
	}

	header["Content-Type"] = "application/json;charset=UTF-8"

	switch method {
	case http.MethodGet:
		if jsonData, resp, err = getDataWithHeader(uri, header); err != nil {
			return err
		}
	case http.MethodPost:
		if jsonData, resp, err = postDataWithHeader(uri, p.Body(), header); err != nil {
			return err
		}
	case http.MethodPut:
		if jsonData, resp, err = putDataWithHeader(uri, p.Body(), header); err != nil {
			return err
		}
	case http.MethodDelete:
		if jsonData, resp, err = deleteDataWithHeader(uri, header); err != nil {
			return err
		}
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	default:
		non200err, err := resolveNon2XX(resp, jsonData)
		if err != nil {
			return err
		}
		return wrapWithAPIErr(non200err)
	}
	return jsoniter.Unmarshal(jsonData, i)
}

func resolveNon2XX(res *http.Response, body []byte) (*resources.Non2XXError, error) {
	var err error
	non200err := &resources.Non2XXError{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}

	switch util.HeaderValues("Content-Type", res.Header) {
	case "":
		non200err.APIErrors = []resources.ErrorInformation{
			{Message: "Content-Type is undefined."},
		}
		return non200err, nil
	case "application/json":
		if body == nil {
			if body, err = io.ReadAll(res.Body); err != nil {
				return nil, err
			}
		}

		if err := jsoniter.Unmarshal(body, non200err); err != nil {
			return nil, err
		}
	default:
		if body == nil {
			if body, err = io.ReadAll(res.Body); err != nil {
				return nil, err
			}
		}
		non200err.APIErrors = []resources.ErrorInformation{
			{Message: strings.TrimRight(string(body), "\n")},
		}
		return non200err, nil
	}

	// additional information for Rate Limit
	if res.StatusCode == http.StatusTooManyRequests {
		if non200err.RateLimitInfo, err = util.GetRateLimitInformation(res); err != nil {
			return nil, err
		}
	}

	return non200err, nil
}
