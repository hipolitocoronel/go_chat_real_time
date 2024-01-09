package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"password"`
}
