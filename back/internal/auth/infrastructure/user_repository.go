package infrastructure

import (
	"go_real_time_chat/config/database"
	"go_real_time_chat/internal/auth/domain"
)

type UserRepository interface {
	FindByEmail(email string) *domain.User
	CreateUser(user *domain.User) (*domain.User, error)
}

type userRepository struct {
	DB *database.MySQLClient
}

func NewUserRepository(db *database.MySQLClient) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) FindByEmail(email string) *domain.User {
	return &domain.User{}
}


func (ur *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}