package infrastructure

import (
	"go_real_time_chat/config/database"
	"go_real_time_chat/internal/auth/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userRepository struct {
	DB *database.MySQLClient
}

func NewUserRepository(db *database.MySQLClient) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := ur.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		//si el error es distinto a que no haya registros -> retornar error
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	return &user, nil
}

// ----CREAR USUARIO
func (ur *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ----OBETENER USUARIOS
func (ur *userRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	result := ur.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
