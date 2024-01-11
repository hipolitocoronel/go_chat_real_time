package database

import (
	"fmt"
	"go_real_time_chat/internal/auth/domain"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	*gorm.DB
}

func InitMySQL() (*MySQLClient, error) {
	// 1. get environment variables
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	// 2. dsn config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	// 3. make connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Print("Connection with database successfull")

	//crear tablas
	db.AutoMigrate(&domain.User{})

	return &MySQLClient{db}, nil
}
