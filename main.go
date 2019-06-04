package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/gin-gonic/gin"
)

var credential *oauth.Credentials
var api *anaconda.TwitterApi

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	return anaconda.NewTwitterApi(
		"",
		"",
	)
}

func main() {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		api = getTwitterApi()
		url, cred, err := api.AuthorizationURL("http://mstn2050.com/twitter/callback")
		if err != nil {
			panic(err)
		}
		credential = cred
		c.Redirect(http.StatusMovedPermanently, url)
	})

	g.GET("/twitter/callback", func(c *gin.Context) {
		oauth_verifier := c.Query("oauth_verifier")
		cli, _, err := api.GetCredentials(credential, oauth_verifier)
		if err != nil {
			panic(err)
		}
		api = anaconda.NewTwitterApi(cli.Token, cli.Secret)

		v := url.Values{}
		v.Set("count", "10")

		tweets, err := api.GetHomeTimeline(v)
		if err != nil {
			panic(err)
		}

		for _, tweet := range tweets {
			fmt.Printf("%v\n\n", tweet.FullText)
		}
	})

	g.Run(":80")
}
