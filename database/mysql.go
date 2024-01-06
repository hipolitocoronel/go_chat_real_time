package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_chat"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
