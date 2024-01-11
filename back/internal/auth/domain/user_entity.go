package domain

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Username    string    `json:"username" gorm:"unique"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
