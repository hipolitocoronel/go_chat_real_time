package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"go_real_time_chat/config/oauth"
	"go_real_time_chat/internal/auth/domain"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthService interface {
	GetUserInfo(token *oauth2.Token) (string, error)
}

type GoogleOAuthService struct {
	AuthConf oauth.OAuthConf
}

func (s *GoogleOAuthService) GetUserInfo(token *oauth2.Token) (domain.UserGoogle, error) {
	// get user info from google api
	userInfoURL := "https://www.googleapis.com/oauth2/v3/userinfo"
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	var userInfo domain.UserGoogle

	resp, err := client.Get(userInfoURL)
	if err != nil {
		return userInfo, errors.New("Error getting user info from google")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return userInfo, errors.New("Bad response from google server: " + resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return userInfo, errors.New("Error while decoding google user: " + err.Error())
	}

	return userInfo, nil
}
