package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("KsWKFfrrKnNfrNms7ZE7VEzVU")
	anaconda.SetConsumerSecret("a6hKnH3rTNUaOim4muDrScgKJ76oAzzI90VBXKo5Ejssu5lBcs")
	return anaconda.NewTwitterApi(
		"",
		"",
	)
}

func main() {

	api := getTwitterApi()

	uri, cred, err := api.AuthorizationURL("oob")
	if err != nil {
		panic(err)
	}

	fmt.Print(uri)
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	oauth_verifier := stdin.Text()

	cred, _, err = api.GetCredentials(cred, oauth_verifier)
	if err != nil {
		panic(err)
	}

	api = anaconda.NewTwitterApi(cred.Token, cred.Secret)
	fmt.Println("authorize successful")

	text := "Hello world"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)
}
