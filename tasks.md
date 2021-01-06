# Tasks

1. Create mail account:
twitchplays.jd@gmail.com : jnkdtzl264

2. Create twitch account:
twitchplaysjd : jnkdtzl264

3. Create twitch application:   
https://dev.twitch.tv/console/apps
*(Add 2-Factor-Authentication)*
Name: chatplays
OAuth Redirect URLs: http://localhost:8080

4. Copy Client ID & Client Secret:
Client ID: 0a4o4z4ivddgafxq76u4oq9c847978
Client Secret: k7a2qzi9ojo7ulhfo0hksxq6dd2rap

5. Setup OAuth Server:
- clone go implementation from twitchs github repo
  https://github.com/twitchdev/authentication-go-sample
- install dependencies
  "go get -u github.com/twitchdev/authentication-go-sample/..."
- change client_id, client_secret and redirect_url in each (main.go) file
- change the scope of information to 
"scopes = []string{"chat:read", "chat:edit"}"
- start the server 
cd into oauth-authorization-code and execute "go run main.go"
-  build an executable
"go build"
- go to localhost:8080 and login with your twitch account

6. Setup IRC Client:


AccessToken: cupb1bl3y50pgnx3bc7l4rv690efs3
Muss ich jedes Mal einen neuen Token genereieren?
Brauche ich evtl. die anderen beiden OAuth Dateien?

Warum klappt das mit dem IRC nicht?
Gibt es noch ein einfacheres Beispiel?




