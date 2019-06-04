package main

import (
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/gin-gonic/gin"
)

var credential *oauth.Credentials
var api *anaconda.TwitterApi

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("KsWKFfrrKnNfrNms7ZE7VEzVU")
	anaconda.SetConsumerSecret("a6hKnH3rTNUaOim4muDrScgKJ76oAzzI90VBXKo5Ejssu5lBcs")
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
		api.PostTweet("眠い", nil)
	})

	g.Run(":80")
}
