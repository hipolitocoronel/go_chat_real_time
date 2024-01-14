package handler

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/domain"
	"go_real_time_chat/internal/auth/service"
	"go_real_time_chat/pkg/dtos/userdtos"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(config *config.AppConfig, router fiber.Router, us *service.UserService) {
	// auth
	router.Post("/login", Login(us))
	router.Get("/login/google", GoogleLogin(config))

	//users
	router.Post("/users", CreateUser(us))
	router.Get("/users", GetUsers(us))

}

func Login(us *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userReq userdtos.UserLoginReq

		if err := c.BodyParser(&userReq); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status": false,
				"error":  "body parser: " + err.Error(),
			})
		}

		// case success
		return nil
	}
}

func GoogleLogin(config *config.AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := config.AuthConf.AuthCodeURL("randomstate")

		c.Status(fiber.StatusSeeOther)
		// return c.Redirect(url)
		return c.JSON(url)
	}
}

func CreateUser(us *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user domain.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		createdUser, err := us.CreateUserService(&user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(createdUser)
	}
}

func GetUsers(us *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := us.GetUsersService()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "users retrieved successfully",
			"status":  fiber.StatusOK,
			"data":    users,
		})
	}
}