package handler

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(config *config.AppConfig, router fiber.Router, userUseCase *usecase.AuthUseCase) {
	router.Get("/google-login", func(c *fiber.Ctx) error {
		return GoogleLogin(c, config)
	})
}

func GoogleLogin(c *fiber.Ctx, config *config.AppConfig) error {
	url := config.AuthConf.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}
