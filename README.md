# TwitchPlays
TwitchPlays Bot written in Golang. Let the viewers play your game via chat commands.

# Setup 
The Bot consists of 2 modules. The OAuth and the IRC-Client.
First, we need to authenticate our bot and collect the access-token.
After that, we can start the IRC-Client by providing the token.

# Customizing
The available chat commands and their specific implementation is realised in the IRC-Client.
Check and modify this file as you want.
