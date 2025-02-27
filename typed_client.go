package gotwitter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/928799934/go-twitter/internal/util"
)

type TypedClient[T util.Response] struct {
	c *GoTwitter
}

func NewTypedClient[T util.Response](c *GoTwitter) *TypedClient[T] {
	if c == nil {
		return nil
	}

	return &TypedClient[T]{
		c: c,
	}
}

func (c *TypedClient[T]) CallStream(ctx context.Context, uri string, method HTTPMethod, auth AuthMethod, p util.Parameters) (*StreamClient[T], error) {

	var (
		header map[string]string
		err    error
		resp   *http.Response
	)

	uri = p.ResolveEndpoint(uri)

	switch auth {
	case OAuth1UserContext:
		params := url.Values{}
		for k, v := range p.ParameterMap() {
			params.Set(k, v)
		}

		if header, err = getOAuth1Header(c.c, method, uri, params); err != nil {
			return nil, err
		}
	case OAuth2BearerToken:
		if header, err = getOAuth2BearerTokenHeader(c.c); err != nil {
			return nil, err
		}
	case OAuth2AccessToken:
		if header, err = getOAuth2AccessTokenHeader(c.c); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("AuthMethod not support")
	}

	header["Content-Type"] = "application/json;charset=UTF-8"

	switch method {
	case http.MethodGet:
		if resp, err = doDataWithHeader(uri, method, nil, header); err != nil {
			return nil, err
		}
	case http.MethodPost:
		if resp, err = doDataWithHeader(uri, method, p.Body(), header); err != nil {
			return nil, err
		}
	case http.MethodPut:
		if resp, err = doDataWithHeader(uri, method, p.Body(), header); err != nil {
			return nil, err
		}
	case http.MethodDelete:
		if resp, err = doDataWithHeader(uri, method, nil, header); err != nil {
			return nil, err
		}
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	default:
		non200err, err := resolveNon2XX(resp, nil)
		if err != nil {
			return nil, err
		}
		return nil, wrapWithAPIErr(non200err)
	}
	return newStreamClient[T](resp)
}
