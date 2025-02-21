package gotwitter

import (
	"crypto/tls"
	"io"
	"net/http"
	"regexp"
	"time"
)

var (
	readTimeout = 30 * time.Second

	// + => %20
	regexpQuery = regexp.MustCompile(`([^%])(\+)`)
)

func getTransport() http.RoundTripper {
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	tr.DisableCompression = true
	return tr
}

func getDataWithHeader(url string /*params url.Values,*/, headers map[string]string) ([]byte, *http.Response, error) {

	/*
		buf := bytes.NewBufferString(url)
		if params != nil {
			buf.WriteString("?")

			// " " -> "+"  must be "%20"
			buf.WriteString(regexpQuery.ReplaceAllString(params.Encode(), "$1%20"))
		}
		resp, err := doDataWithHeader(buf.String(), http.MethodGet, nil, headers)
	*/

	resp, err := doDataWithHeader(url, http.MethodGet, nil, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func postDataWithHeader(url string, ioParams io.Reader, headers map[string]string) ([]byte, *http.Response, error) {
	resp, err := doDataWithHeader(url, http.MethodPost, ioParams, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func putDataWithHeader(url string, ioParams io.Reader, headers map[string]string) ([]byte, *http.Response, error) {
	resp, err := doDataWithHeader(url, http.MethodPut, ioParams, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func deleteDataWithHeader(url string, headers map[string]string) ([]byte, *http.Response, error) {
	resp, err := doDataWithHeader(url, http.MethodDelete, nil, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func doDataWithHeader(url, method string, ioParams io.Reader, headers map[string]string) (*http.Response, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}

	req, err := http.NewRequest(method, url, ioParams)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return cli.Do(req)
}
