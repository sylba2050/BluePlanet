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

	for range time.Tick(time.Minute * 1) {
		rand.Seed(time.Now().Unix())
		num := rand.Intn(3) + 3
		var tweet string
		for len([]rune(tweet)) < num {
			r := randomChoice("あいうえおかきくけこさしすせそたちつてとなにぬのはひふへほまみむめもやゆよらりるれろわゐをゑんぁぃぅぇぉゃゅょゎ")
			if !isInclude(tweet, r) {
				tweet += string(r)
			}
		}

		client := twitter.NewClient(httpClient)
		_, _, err := client.Statuses.Update(tweet, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("done")
	}
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
