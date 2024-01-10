package oauth

import (
	"fmt"
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
	APP_HOST := os.Getenv("APP_HOST")
	APP_PORT := os.Getenv("APP_PORT")
	APP_VERSION := os.Getenv("APP_VERSION")

	// 2. initialize configuration
	config := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  fmt.Sprintf("%s:%s/%s/auth/google-callback", APP_HOST, APP_PORT, APP_VERSION), // TODO: get from environmet
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint}

	// return auth config
	return &OAuthConf{config}
}
