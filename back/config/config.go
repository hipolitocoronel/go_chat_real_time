package config

import (
	"fmt"
	"go_real_time_chat/config/database"
	oauth2_conf "go_real_time_chat/config/oauth2"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB         *gorm.DB
	GoogleAuth *oauth2.Config
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

	fmt.Print("Connection with database successfull")

	// 3. get google auth conf
	googleConf := oauth2_conf.InitGoogleAuth()

	// return app conf
	return &AppConfig{
		DB:         db,
		GoogleAuth: googleConf,
	}, nil
}
