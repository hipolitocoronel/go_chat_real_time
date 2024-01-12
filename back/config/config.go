package config

import (
	"go_real_time_chat/config/database"
	"go_real_time_chat/config/oauth"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB       *database.MySQLClient
	AuthConf *oauth.OAuthConf
	Cors     cors.Config
	Port     string
}

func LoadConfig() (*AppConfig, error) {
	// 1. load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return nil, err
	}

	// 2. connect to database
	db, err := database.InitMySQL()

	if err != nil {
		log.Fatal("Error connecting database: ", err)
		return nil, err
	}

	// 3. get google auth conf
	authConf := oauth.InitOAuthConf()

	// 4. config cors
	cors := cors.Config{
		AllowOrigins: "*",
	}

	// 5. config port
	port := ":" + os.Getenv("APP_PORT")
	if len(port) == 1 {
		port += "3000" // default port
	}

	// return app conf
	return &AppConfig{
		DB:       db,
		AuthConf: authConf,
		Cors:     cors,
		Port:     port,
	}, nil
}
