package main

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/handler"
	"go_real_time_chat/internal/auth/infrastructure"
	"go_real_time_chat/internal/auth/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app := fiber.New()
	app.Use(cors.New(config.Cors))
	versionApp := app.Group("/api/v1")

	// user
	userRepository := infrastructure.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)

	handler.UserRoutes(config, versionApp, userService)

	app.Listen(config.Port)
}
