package domain

import (
	"errors"
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

// validar que el usuario tenga todos los campos
func (user *User) Validate() error {

	if len(user.Email) == 0 {
		return errors.New("email is required")
	}

	if len(user.Password) == 0 {
		return errors.New("password is required")
	}

	if len(user.Name) == 0 {
		return errors.New("name is required")
	}

	if len(user.Username) == 0 {
		return errors.New("username is required")
	}

	if len(user.PhoneNumber) == 0 {
		return errors.New("phone number is required")
	}

	return nil
}
