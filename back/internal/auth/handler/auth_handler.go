package handler

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/domain"
	"go_real_time_chat/internal/auth/usecase"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(config *config.AppConfig, router fiber.Router, userUseCase *usecase.AuthUseCase) {
	router.Get("/google-login", func(c *fiber.Ctx) error {
		return GoogleLogin(c, config)
	})

	router.Get("/google-callback", func(c *fiber.Ctx) error {
		return GoogleCallback(c)
	})

	//users
	router.Post("/users", func(c *fiber.Ctx) error {
		return CreateUser(c, userUseCase)
	})

	router.Get("/users", func(c *fiber.Ctx) error {
		return GetUsers(c, userUseCase)
	})

}

func GoogleLogin(c *fiber.Ctx, config *config.AppConfig) error {
	url := config.AuthConf.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	// return c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	FRONT_URL := os.Getenv("APP_FRONT_URL")
	code := c.FormValue("code")

	cookie := &fiber.Cookie{
		Name:     "code",
		Value:    code,
		SameSite: "None",
		Domain:   FRONT_URL,
		Path:     FRONT_URL,
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	c.Cookie(cookie)
	c.Context().Redirect(FRONT_URL, 200)
	return c.JSON(200, "hola")
}

func CreateUser(c *fiber.Ctx, userUseCase *usecase.AuthUseCase) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdUser, err := userUseCase.CreateUserService(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdUser)
}

func GetUsers(c *fiber.Ctx, userUseCase *usecase.AuthUseCase) error {

	users, err := userUseCase.GetUsersService()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "users retrieved successfully",
		"status":  fiber.StatusOK,
		"data":    users,
	})
}
