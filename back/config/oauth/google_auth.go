package oauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthConf struct {
	*oauth2.Config
}

func InitOAuthConf() *OAuthConf {
	// 1. get environment variables
	CLIENT_ID := os.Getenv("AUTH_GOOGLE_ID")
	CLIENT_SECRET := os.Getenv("AUTH_GOOGLE_SECRET")

	// 2. initialize configuration
	config := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  "http://127.0.0.1:3000/callback", // TODO: get from environmet
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint}

	// return auth config
	return &OAuthConf{config}
}
