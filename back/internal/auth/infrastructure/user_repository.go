package infrastructure

import (
	"go_real_time_chat/config/database"
	"go_real_time_chat/internal/auth/domain"
)

type UserRepository interface {
	FindByEmail(email string) *domain.User
}

type userRepository struct {
	DB *database.MySQLClient
}
