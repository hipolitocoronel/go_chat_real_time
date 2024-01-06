package main

import (
	"fmt"
	"go_real_time_chat/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.SetupDB()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(db)

	app := fiber.New()

	app.Listen(":3000")
}
