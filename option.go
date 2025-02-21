package gotwitter

type OptionFunc func(*GoTwitter)

func WithOAuth(apiKey, apiKeySecret, accessToken, accessTokenSecret string) OptionFunc {
	return func(t *GoTwitter) {
		t.apiKey = apiKey
		t.apiKeySecret = apiKeySecret
		t.accessToken = accessToken
		t.accessTokenSecret = accessTokenSecret
	}
}

func WithBearerToken(bearerToken string) OptionFunc {
	return func(t *GoTwitter) {
		t.bearerToken = bearerToken
	}
}

// OAuth 2.0 Client ID and Client Secret 未使用
// func WithOAuth2(clientID, clientSecret string) OptionFunc {
// 	return func(t *GoTwitter) {
// 		t.clientID = clientID
// 		t.clientSecret = clientSecret
// 	}
// }
