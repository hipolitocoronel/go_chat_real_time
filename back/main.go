package main

import (
	"fmt"
	"go_real_time_chat/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// cargando variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// realizando conexión a la DB
	db, err := database.InitMySQL()

	if err != nil {
		log.Fatal("Error al conectar la base de datos: ", err)
	}

	fmt.Print("Conexión exitosa", db)

	app := fiber.New()

	app.Listen(":3000")
}
