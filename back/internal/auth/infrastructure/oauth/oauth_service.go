package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"go_real_time_chat/config/oauth"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthService interface {
	GetUserInfo(token *oauth2.Token) (string, error)
}

type GoogleOAuthService struct {
	AuthConf oauth.OAuthConf
}

func (s *GoogleOAuthService) GetUserInfo(token *oauth2.Token) (string, error) {
	// get user info from google api
	userInfoURL := "https://www.googleapis.com/oauth2/v3/userinfo"
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	resp, err := client.Get(userInfoURL)
	if err != nil {
		return "", fmt.Errorf("error al obtener información del usuario desde Google: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("respuesta no exitosa del servidor de Google: %s", resp.Status)
	}

	var userInfo struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", fmt.Errorf("error al decodificar la respuesta JSON de Google: %v", err)
	}

	return userInfo.Email, nil
}
