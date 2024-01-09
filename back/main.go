package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_real_time_chat/config"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserInfo struct {
	Name     string `json:"name"`
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	Verified bool   `json:"email_verified"`
	Picture  string `json:"picture"`
}

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Get("/login", loginHandler)
	app.Get("/callback", callbackHandler)

	app.Listen(":3000")
}

func loginHandler(c *fiber.Ctx) error {
	config := &oauth2.Config{
		ClientID:     os.Getenv("AUTH_GOOGLE_ID"),
		ClientSecret: os.Getenv("AUTH_GOOGLE_SECRET"),
		RedirectURL:  "http://127.0.0.1:3000/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	url := config.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func callbackHandler(c *fiber.Ctx) error {
	config := &oauth2.Config{
		ClientID:     os.Getenv("AUTH_GOOGLE_ID"),
		ClientSecret: os.Getenv("AUTH_GOOGLE_SECRET"),
		RedirectURL:  "http://127.0.0.1:3000/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	code := c.Query("code")

	// Exchange the authorization code for a token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Token exchange error: %v", err))
	}

	if !token.Valid() {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Token not valid: %v", err))
	}

	// Retrieve user information using the obtained token
	userInfo, err := fetchGoogleUserInfo(token.AccessToken)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error fetching Google user info: %v", err))
	}

	// Perform any additional actions with the user information (e.g., store in the database)

	return c.SendString(fmt.Sprintf("Login successful! User ID: %s, Email: %s", userInfo.Sub, userInfo.Email))
}

func fetchGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, err
	}

	// Agregar el token de acceso al encabezado de autorización
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Respuesta no exitosa de Google. Código de estado: %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
