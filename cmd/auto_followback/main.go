package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
	)
	token := oauth1.NewToken(
		os.Getenv("RANDOM_BOT_KEY"),
		os.Getenv("RANDOM_BOT_SECRET"),
	)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	for range time.Tick(time.Minute * 1) {
		followers := twitter.FollowerIDParams{
			ScreenName: "randomb24532179",
		}
		followersID, _, _ := client.Followers.IDs(&followers)

		follows := twitter.FriendIDParams{
			ScreenName: "randomb24532179",
		}
		followsID, _, _ := client.Friends.IDs(&follows)

		notFollowing := sliceSub(followersID.IDs, followsID.IDs)
		notFollower := sliceSub(followsID.IDs, followersID.IDs)

		fmt.Println(notFollowing)
		fmt.Println(notFollower)
		for _, v := range notFollowing {
			followParam := twitter.FriendshipCreateParams{
				UserID: v,
			}
			client.Friendships.Create(&followParam)
		}

		for _, v := range notFollower {
			unfollowParam := twitter.FriendshipDestroyParams{
				UserID: v,
			}
			client.Friendships.Destroy(&unfollowParam)
		}
	}
}

func in(target []int64, element int64) bool {
	for _, v := range target {
		if v == element {
			return true
		}
	}
	return false
}

func sliceSub(a, b []int64) []int64 {
	var res []int64
	for _, v := range a {
		if !in(b, v) {
			res = append(res, v)
		}
	}
	return res
}
