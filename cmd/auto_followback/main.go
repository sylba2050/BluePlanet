package main

import (
	"fmt"
	"math/rand"
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

	followers := twitter.FollowerIDParams{
		ScreenName: "randomb24532179",
	}
	followersID, followersRes, _ := client.Followers.IDs(&followers)
	fmt.Println(followersID.IDs)
	fmt.Println(followersRes)

	follows := twitter.FriendIDParams{
		ScreenName: "randomb24532179",
	}
	followsID, followsRes, _ := client.Friends.IDs(&follows)
	fmt.Println(followsID.IDs)
	fmt.Println(followsRes)

	followParam := twitter.FriendshipCreateParams{
		UserID: followersID.IDs[0],
	}
	client.Friendships.Create(&followParam)

	unfollowParam := twitter.FriendshipDestroyParams{
		UserID: followersID.IDs[0],
	}
	client.Friendships.Destroy(&unfollowParam)
}

func randomChoice(from string) rune {
	runes := []rune(from)

	rand.Seed(time.Now().Unix())
	return runes[rand.Intn(len(runes))]
}

func isInclude(target string, element rune) bool {
	runes := []rune(target)
	for _, v := range runes {
		if v == element {
			return true
		}
	}

	return false
}
