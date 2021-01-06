package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc"
)

func main() {
	token := "9di4v38hkd42qi180tpm2kzmorn1ta"

	client := twitch.NewClient("twitchplaysjd", fmt.Sprintf("oauth:%s", token))

	client.OnPrivateMessage(
		func(message twitch.PrivateMessage) {

			fmt.Printf("%s said: %s", message.User.DisplayName, message.Message)
			client.Say(message.Channel, message.User.DisplayName+" someone said: "+message.Message)
		},
	)

	client.Join("twitchplaysjd")

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
