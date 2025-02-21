gotwitter
===

This is a library for using the Twitter API v2 in the Go language. (It is still under development).


**The code is heavily modified from [https://github.com/michimani/gotwi](https://github.com/michimani/gotwi).**


# Supported APIs

[Twitter API Documentation | Docs | Twitter Developer Platform  ](https://developer.twitter.com/en/docs/twitter-api)

Progress of supporting APIs:

| Category | Sub Category | Endpoint |
| --- | --- | --- |
| Tweets | Tweet lookup | `GET /2/tweets` |
|  |  | `GET /2/tweets/:id` |
|  | Manage Tweet | `POST /2/tweets` |
|  |  | `DELETE /2/tweets/:id` |
|  | Timelines | `GET /2/users/:id/tweets` |
|  |  | `GET /2/users/:id/mentions` |
|  |  | `GET /2/users/:id/timelines/reverse_chronological` |
|  | Search Tweets | `GET /2/tweets/search/recent` |
|  |  | `GET /2/tweets/search/all` |
|  | Tweet counts | `GET /2/tweets/counts/recent` |
|  |  | `GET /2/tweets/counts/all` |
|  | Filtered stream | `POST /2/tweets/search/stream/rules` |
|  |  | `GET /2/tweets/search/stream/rules` |
|  |  | `GET /2/tweets/search/stream` |
|  | Volume streams | `GET /2/tweets/sample/stream` |
|  | Retweets | `GET /2/users/:id/retweeted_by` |
|  |  | `POST /2/users/:id/retweets` |
|  |  | `DELETE /2/users/:id/retweets/:source_tweet_id` |
|  | Likes | `GET /2/tweets/:id/liking_users` |
|  |  | `GET /2/tweets/:id/liked_tweets` |
|  |  | `POST /2/users/:id/likes` |
|  |  | `DELETE /2/users/:id/likes/:tweet_id` |
|  | Hide replies | `PUT /2/tweets/:id/hidden` |
|  | Quote Tweets | `GET /2/tweets/:id/quote_tweets` |
|  | Bookmarks | `GET /2/users/:id/bookmarks` |
|  |  | `POST /2/users/:id/bookmarks` |
|  |  | `DELETE /2/users/:id/bookmarks/:tweet_id` |
| Users | User lookup | `GET /2/users` |
|  |  | `GET /2/users/:id` |
|  |  | `GET /2/users/by` |
|  |  | `GET /2/users/by/username` |
|  |  | `GET /2/users/by/me` |
|  | Follows | `GET /2/users/:id/following` |
|  |  | `GET /2/users/:id/followers` |
|  |  | `POST /2/users/:id/following` |
|  |  | `DELETE /2/users/:source_user_id/following/:target_user_id` |
|  | Blocks | `GET /2/users/:id/blocking` |
|  |  | `POST /2/users/:id/blocking` |
|  |  | `DELETE /2/users/:source_user_id/blocking/:target_user_id` |
|  | Mutes | `GET /2/users/:id/muting` |
|  |  | `POST /2/users/:id/muting` |
|  |  | `DELETE /2/users/:source_user_id/muting/:target_user_id` |
| Lists | List lookup | `GET /2/lists/:id` |
|  |  | `GET /2/users/:id/owned_lists` |
|  | Manage Lists | `POST /2/lists` |
|  |  | `DELETE /2/lists/:id` |
|  |  | `PUT /2/lists/:id` |
|  | List Tweets lookup | `GET /2/lists/:id/tweets` |
|  | List members | `GET /2/users/:id/list_memberships` |
|  |  | `GET /2/lists/:id/members` |
|  |  | `POST /2/lists/:id/members` |
|  |  | `DELETE /2/lists/:id/members/:user_id` |
|  | List follows | `GET /2/lists/:id/followers` |
|  |  | `GET /2/users/:id/followed_lists` |
|  |  | `POST /2/users/:id/followed_lists` |
|  |  | `DELETE /2/users/:id/followed_lists/:list_id` |
|  | Pinned Lists | `GET /2/users/:id/pinned_lists` |
|  |  | `POST /2/users/:id/pinned_lists` |
|  |  | `DELETE /2/users/:id/pinned_lists/:list_id` |
| Spaces | Spaces Lookup | `GET /2/spaces/:id` |
|  |  | `GET /2/spaces` |
|  |  | `GET /2/spaces/by/creator_ids` |
|  |  | `GET /2/spaces/:id/buyers` |
|  |  | `GET /2/spaces/:id/tweets` |
|  | Search Spaces | `GET /2/spaces/search` |
| Compliance | Batch compliance | `GET /2/compliance/jobs/:id` |
|  |  | `GET /2/compliance/jobs` |
|  |  | `POST /2/compliance/jobs` |


# How to use

With this authentication method, each operation will be performed as the authenticated Twitter account. For example, you can tweet as that account, or retrieve accounts that are blocked by that account.

### Example: Get your own information.

```go
package main

import (
	"context"
	"fmt"

	"github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/user/userlookup"
	"github.com/928799934/go-twitter/user/userlookup/types"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
	BearerToken       = ""
)

func main() {
	c := gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))

	p := &types.GetMeInput{
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetMe(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:          ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name:        ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username:    ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt:   ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
}
```

```
go run main.go
```

You will get the output like following.

```
ID:           581780917
Name:         michimani Lv.873
Username:     michimani210
CreatedAt:    2012-05-16 12:07:04 +0000 UTC
PinnedTweet:  OpenAI API の Function Calling を使って自然言語で AWS リソースを作成してみる
```

### Example: Tweet with poll.

```go
package main

import (
	"context"
	"fmt"

	"github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/tweet/managetweet"
	"github.com/928799934/go-twitter/tweet/managetweet/types"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
)

func main() {
	c := gotwitter.NewGoTwitter(gotwitter.WithOAuth1(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))

	p := &types.CreateInput{
		Text: gotwi.String("This is a test tweet with poll."),
		Poll: &types.CreateInputPoll{
			DurationMinutes: gotwi.Int(5),
			Options: []string{
				"Cyan",
				"Magenta",
				"Yellow",
				"Key plate",
			},
		},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
```

```
go run main.go
```

You will get the output like following.

```
[1462813519607263236] This is a test tweet with poll.
```

## Request with OAuth 2.0 Bearer Token

This authentication method allows only read-only access to public information.

### Example: Get a user by user name.

⚠ This example only works with Twitter API v2 Basic or Pro plan. see details: [Developers Portal](https://developer.twitter.com/en/portal/products)

```go
package main

import (
	"context"
	"fmt"

	"github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/fields"
	"github.com/928799934/go-twitter/user/userlookup"
	"github.com/928799934/go-twitter/user/userlookup/types"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
	BearerToken       = "" // If a value is present, it must be the correct value.
)

func main() {
	c := gotwitter.NewGoTwitter(gotwitter.WithOAuth1(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret),gotwitter.WithBearerToken(BearerToken))
	// OR
	// c := gotwitter.NewGoTwitter(gotwitter.WithOAuth1(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))

	p := &types.GetByUsernameInput{
		Username: "michimani210",
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetByUsername(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:          ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name:        ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username:    ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt:   ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
}
```

```
go run main.go
```

You will get the output like following.

```
ID:           581780917
Name:         michimani Lv.861
Username:     michimani210
CreatedAt:    2012-05-16 12:07:04 +0000 UTC
PinnedTweet:  真偽をハッキリしたい西城秀樹「ブーリアン、ブーリアン」
```

## Error handling

Each function that calls the Twitter API (e.g. `retweet.ListUsers()`) may return an error for some reason.
If the error is caused by the Twitter API returning a status other than 2XX, you can check the details by doing the following.

```go
res, err := retweet.ListUsers(context.Background(), c, p)
if err != nil {
	fmt.Println(err)

	// more error information
	ge := err.(*gotwi.GotwiError)
	if ge.OnAPI {
		fmt.Println(ge.Title)
		fmt.Println(ge.Detail)
		fmt.Println(ge.Type)
		fmt.Println(ge.Status)
		fmt.Println(ge.StatusCode)

		for _, ae := range ge.APIErrors {
			fmt.Println(ae.Message)
			fmt.Println(ae.Label)
			fmt.Println(ae.Parameters)
			fmt.Println(ae.Code)
			fmt.Println(ae.Code.Detail())
		}

		if ge.RateLimitInfo != nil {
			fmt.Println(ge.RateLimitInfo.Limit)
			fmt.Println(ge.RateLimitInfo.Remaining)
			fmt.Println(ge.RateLimitInfo.ResetAt)
		}
	}
}
```


## More examples

See [_examples](https://github.com/928799934/go-twitter/tree/main/_examples) directory.

# Licence

[MIT](https://github.com/928799934/go-twitter/blob/main/LICENCE)

# Author

[928799934](https://github.com/928799934)
