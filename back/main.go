package main

import (
	"go_real_time_chat/config"
	"go_real_time_chat/internal/auth/handler"
	"go_real_time_chat/internal/auth/infrastructure"
	"go_real_time_chat/internal/auth/usecase"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("")
		return
	}

	app := fiber.New()
	versionApp := app.Group("/api/v1")

	// auth
	authRouter := versionApp.Group("/auth")
	userRepository := infrastructure.NewUserRepository(config.DB)
	userUseCase := usecase.NewAuthUseCase(userRepository)

	handler.AuthRoutes(config, authRouter, userUseCase)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
