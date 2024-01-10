package handler

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/usecase"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(config *config.AppConfig, router fiber.Router, userUseCase *usecase.AuthUseCase) {
	router.Get("/google-login", func(c *fiber.Ctx) error {
		return GoogleLogin(c, config)
	})

	router.Get("/google-callback", func(c *fiber.Ctx) error {
		return GoogleCallback(c)
	})
}

func GoogleLogin(c *fiber.Ctx, config *config.AppConfig) error {
	url := config.AuthConf.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	FRONT_URL := os.Getenv("APP_FRONT_URL")
	code := c.Query("code")

	c.Redirect(FRONT_URL)
	return c.JSON(code)
}
