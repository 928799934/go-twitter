package util

import (
	"io"
	"net/url"
	"regexp"
	"strings"
)

var (
	// + => %20
	regexpQuery = regexp.MustCompile(`([^%])(\+)`)
)

type Parameters interface {
	ResolveEndpoint(endpointBase string) string
	Body() io.Reader
	ContentType() string
	ParameterMap() map[string]string
}

func QueryValue(params []string) string {
	if len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}

func QueryString(paramsMap map[string]string, includes map[string]struct{}) string {
	q := url.Values{}
	for k, v := range paramsMap {
		if _, ok := includes[k]; ok {
			q.Add(k, v)
		}
	}

	return regexpQuery.ReplaceAllString(q.Encode(), "$1%20")
}
