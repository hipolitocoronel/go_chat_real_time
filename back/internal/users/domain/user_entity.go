package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	PhoneNumber string
	Email       string `gorm:"unique;not null"`
	Password    string
}
