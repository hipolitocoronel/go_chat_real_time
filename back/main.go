package main

import (
	"go_real_time_chat/config"

	"github.com/gofiber/fiber/v2"
)

// type GoogleUserInfo struct {
// 	Name     string `json:"name"`
// 	Sub      string `json:"sub"`
// 	Email    string `json:"email"`
// 	Verified bool   `json:"email_verified"`
// 	Picture  string `json:"picture"`
// }

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Listen(":3000")
}
