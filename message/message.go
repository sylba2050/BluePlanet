package main

import (
	"fmt"
	"log"
	// "github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig("JCL6KkhZO2dHVViWnKaUm6UzJ", "4cEUT6DPlpUWVZcQLqxdGCmDQNFNJZb6U8dGZMdPDQd1ycwhCi")
	token := oauth1.NewToken("1012877640997339136-YBRdMBSxHerPXHSvj32yKEroIqpDNs", "ifjWZ3zyrOCqbmFxfTgZXVFTz2nxjPjNDBS6I7BRNsxGz")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Show User Param
	userShowParams := &twitter.UserShowParams{ScreenName: "Sylba2050"}
	user, _, _ := client.Users.Show(userShowParams)
	fmt.Printf("USERS SHOW:\n%+v\n", user)

	// List most recent 10 Direct Messages
	messages, _, err := client.DirectMessages.EventsList(
		&twitter.DirectMessageEventsListParams{Count: 10},
	)
	fmt.Println("User's DIRECT MESSAGES:")
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range messages.Events {
		fmt.Printf("%+v\n", event)
		fmt.Printf("  %+v\n", event.Message)
		fmt.Printf("  %+v\n", event.Message.Data)
	}

	event, _, err := client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type: "message_create",
			Message: &twitter.DirectMessageEventMessage{
				Target: &twitter.DirectMessageTarget{
					RecipientID: "4834074556",
				},
				Data: &twitter.DirectMessageData{
					Text: "testing",
				},
			},
		},
	})
	fmt.Printf("DM Event New:\n%+v, %v\n", event, err)
}
