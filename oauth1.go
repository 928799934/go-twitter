package gotwitter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

const (
	OAuthVersion10               = "1.0"
	OAuthSignatureMethodHMACSHA1 = "HMAC-SHA1"
)

// OAuth oauth_consumer_key="guozcI4V2Fk4OFsLePpur4Szi",
// oauth_token="1925435862797754368-HTa5eU3etkh6vfFmEXVf1gnn9rzrkO",
// oauth_signature_method="HMAC-SHA1",
// oauth_timestamp="1749380097",
// oauth_nonce="h9vc8BPeUlF",
// oauth_version="1.0",
// oauth_signature="9xffAE3Quc7VutetAThE7xeyhoQ%3D"
func getOAuth1Header(twitter *GoTwitter, method, uri string, params url.Values) (map[string]string, error) {
	if twitter.apiKey == "" || twitter.apiKeySecret == "" || twitter.accessToken == "" || twitter.accessTokenSecret == "" {
		return nil, fmt.Errorf("apiKey, apiKeySecret, accessToken, accessTokenSecret cannot be empty at the same time")
	}
	nonce, err := genOAuthNonce()
	if err != nil {
		return nil, err
	}
	// nonce = "h9vc8BPeUlF"
	timestamp := genOAuthTimestamp()
	// timestamp = "1749380097"

	// http://aaa.com?xxxx=aaaa&xxx=aaaa => http://aaa.com
	uri = getURIBase(uri)

	params.Add("oauth_consumer_key", twitter.apiKey)
	params.Add("oauth_nonce", nonce)
	params.Add("oauth_signature_method", OAuthSignatureMethodHMACSHA1)
	params.Add("oauth_timestamp", timestamp)
	params.Add("oauth_token", twitter.accessToken)
	params.Add("oauth_version", OAuthVersion10)

	parameterString := regexpQuery.ReplaceAllString(params.Encode(), "$1%20")

	sigBase := fmt.Sprintf("%s&%s&%s", url.QueryEscape(method), url.QueryEscape(uri), url.QueryEscape(parameterString))

	signature, err := genOAuthSignature(sigBase, twitter.apiKeySecret, twitter.accessTokenSecret)
	if err != nil {
		return nil, err
	}

	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf(`OAuth oauth_consumer_key="%s",oauth_nonce="%s",oauth_signature="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_token="%s",oauth_version="%s"`,
		url.QueryEscape(twitter.apiKey), url.QueryEscape(nonce), url.QueryEscape(signature), url.QueryEscape(OAuthSignatureMethodHMACSHA1), url.QueryEscape(timestamp), url.QueryEscape(twitter.accessToken), url.QueryEscape(OAuthVersion10))

	return header, nil
}

func genOAuthNonce() (string, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	nonce := base64.StdEncoding.EncodeToString(key)
	symbols := []string{"+", "/", "="}
	for _, s := range symbols {
		nonce = strings.Replace(nonce, s, "", -1)
	}
	return nonce, nil
}

func genOAuthTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func genOAuthSignature(sigBase, apiKeySecret, accessTokenSecret string) (string, error) {
	key := fmt.Sprintf("%s&%s", apiKeySecret, accessTokenSecret)

	h := hmac.New(sha1.New, []byte(key))
	if _, err := io.WriteString(h, sigBase); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func getURIBase(e string) string {
	queryIdx := strings.Index(e, "?")
	if queryIdx < 0 {
		return e
	}

	return e[:queryIdx]
}
