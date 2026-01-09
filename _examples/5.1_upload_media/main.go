package main

import (
	"bytes"
	"fmt"
	"os"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/media/upload/types"
)

const (
	ApiKey            = ""
	ApiKeySecret      = ""
	AccessToken       = ""
	AccessTokenSecret = ""
)

func main() {
	fileBytes, _ := os.ReadFile("sample.png")

	args := os.Args
	isMulti := false
	isPost := false
	if len(args) > 1 && args[1] == "multi" {
		isMulti = true
	}

	if len(args) > 2 && args[2] == "post" {
		isPost = true
	}

	client := newOAuth1Client()

	// Initialize
	res, err := Initialize(client, &types.InitializeInput{
		MediaType:     types.MediaTypePNG,
		TotalBytes:    len(fileBytes),
		Shared:        false,
		MediaCategory: types.MediaCategoryTweetImage,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Initialized media ID: ", res.Data.MediaID)

	// Append
	mediaID := res.Data.MediaID

	if isMulti {
		// Append with some segments
		fmt.Println("Appending with some segments")
		chunkSize := len(fileBytes) / 10
		segmentIndex := 0
		for i := 0; i < len(fileBytes); i += chunkSize {
			end := min(i+chunkSize, len(fileBytes))

			chunk := fileBytes[i:end]
			appendRes, err := Append(client, &types.AppendInput{
				MediaID:      mediaID,
				Media:        bytes.NewReader(chunk),
				SegmentIndex: segmentIndex,
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("Appended segment %d response: %+v\n", segmentIndex, appendRes)
			segmentIndex++
		}
	} else {
		fmt.Println("Appending with single segment")
		appendRes, err := Append(client, &types.AppendInput{
			MediaID:      mediaID,
			Media:        bytes.NewReader(fileBytes),
			SegmentIndex: 0,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("Appended response: %+v\n", appendRes)
	}

	// Finalize
	finalizeRes, err := Finalize(client, &types.FinalizeInput{
		MediaID: mediaID,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Finalized res: %+v\n", finalizeRes)

	// Post media
	if !isPost {
		return
	}

	postedID, err := PostWithMedia(client, "post with a media by using gotwi", mediaID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Posted ID: ", postedID)
}

func newOAuth1Client() *gotwitter.GoTwitter {
	return gotwitter.NewGoTwitter(gotwitter.WithOAuth(ApiKey, ApiKeySecret, AccessToken, AccessTokenSecret))
}
