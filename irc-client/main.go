package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc"
	"github.com/go-vgo/robotgo"
)

func main() {
	token := "cupb1bl3y50pgnx3bc7l4rv690efs3"

	client := twitch.NewClient("twitchplaysjd", fmt.Sprintf("oauth:%s", token))

	client.OnPrivateMessage(
		func(message twitch.PrivateMessage) {

			// Log all chat messages
			fmt.Printf("%s said: %s \n", message.User.DisplayName, message.Message)
			
			// Handle chat commands
			switch message.Message {
			case "a", "A":
				robotgo.KeyTap("q")
			case "b", "B":
				robotgo.KeyTap("w")
			case "start", "Start", "START":
				robotgo.KeyTap("e")
			case "select", "Select", "SELECT":
				robotgo.KeyTap("r")
			case "up", "Up", "UP":
				robotgo.KeyTap("up")
			case "down", "Down", "DOWN": 
				robotgo.KeyTap("down")
			case "left", "Left", "LEFT":
				robotgo.KeyTap("left")
			case "right", "Right", "RIGHT":
				robotgo.KeyTap("right")
			} 
			
		},
	)

	client.Join("twitchplaysjd")

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
