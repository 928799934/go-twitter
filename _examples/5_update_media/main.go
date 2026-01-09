package main

import (
	"context"
	"fmt"
	"os"
	"time"

	gotwitter "github.com/928799934/go-twitter"
	"github.com/928799934/go-twitter/media"
	"github.com/928799934/go-twitter/media/types"
)

var (
	clientAccessToken = "" // 使用clientID 与 clientSecret 经过 oauth2 授权机制获得的accessToken
)

func main() {
	data, _ := os.ReadFile("aoaoao.mp4")

	c := gotwitter.NewGoTwitter(
		gotwitter.WithClientAccessToken(clientAccessToken),
	)

	init, err := media.Init(context.Background(), c, &types.InitInput{
		TotalBytes:    len(data),
		MediaType:     "video/mp4",
		MediaCategory: "tweet_video",
	})
	if err != nil {
		fmt.Println("init", err)
		return
	}
	id := *init.Data.ID
	fmt.Println(id)

	if err := media.Append(context.Background(), c, &types.AppendInput{
		MediaID: id,
		Data:    data,
	}); err != nil {
		fmt.Println("append", err)
		return
	}

	finalize, err := media.Finalize(context.Background(), c, &types.FinalizeInput{
		MediaID: id,
	})
	if err != nil {
		fmt.Println("finalize", err)
		return
	}
	if finalize.Data.ProcessingInfo == nil {
		fmt.Println("update success")
		return
	}

	fmt.Println("status:", *finalize.Data.ProcessingInfo.State)

	// pending -> in_progress -> [failed|succeeded]
	switch *finalize.Data.ProcessingInfo.State {
	case "in_progress":
	case "pending":
	case "succeeded":
		fmt.Println("update success")
		return
	case "failed":
		fmt.Println("update fail")
		return
	}

	for {
		status, err := media.Status(context.Background(), c, &types.StatusInput{
			MediaID: id,
		})
		if err != nil {
			fmt.Println("status", err)
			return
		}
		switch *status.Data.ProcessingInfo.State {
		case "in_progress":
		case "pending":
		case "succeeded":
			fmt.Println("update success with status")
			return
		case "failed":
			fmt.Println("update fail with status")
			return
		}
		time.Sleep(time.Second * 3)
	}
}
