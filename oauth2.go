package gotwitter

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func getOAuth2Header(twitter *GoTwitter) (map[string]string, error) {
	if twitter.bearerToken == "" && (twitter.apiKey == "" || twitter.apiKeySecret == "") {
		return nil, fmt.Errorf("bearerToken, apiKey, apiKeySecret cannot be empty at the same time")
	}

	header := make(map[string]string)

	if twitter.bearerToken == "" {
		params := url.Values{}
		params.Add("grant_type", "client_credentials")

		header["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8"
		header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(twitter.apiKey+":"+twitter.apiKeySecret))

		uri := "https://api.twitter.com/oauth2/token"

		jsonData, resp, err := postDataWithHeader(uri, strings.NewReader(params.Encode()), header)
		if err != nil {
			return nil, err
		}
		delete(header, "Content-Type")
		delete(header, "Authorization")
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code is %d\n message:%s", resp.StatusCode, jsonData)
		}
		var response struct {
			TokenType   string `json:"token_type"`
			AccessToken string `json:"access_token"`
		}
		if err := jsoniter.Unmarshal(jsonData, &response); err != nil {
			return nil, fmt.Errorf("unmarshal error:%s\n message:%s", err.Error(), jsonData)
		}
		if response.AccessToken == "" {
			return nil, fmt.Errorf("access_token is empty")
		}
		twitter.bearerToken = response.AccessToken
	}

	header["Authorization"] = "Bearer " + twitter.bearerToken
	return header, nil
}
