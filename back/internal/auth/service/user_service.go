package service

import (
	"errors"
	"go_real_time_chat/internal/auth/domain"
	"go_real_time_chat/internal/auth/infrastructure"
	"go_real_time_chat/internal/auth/infrastructure/oauth"

	"go_real_time_chat/pkg/dtos/userdtos"

	"golang.org/x/oauth2"
)

type UserService struct {
	UserRepository infrastructure.UserRepository
	OUserService   oauth.OAuthService
}

func NewUserService(userRepo infrastructure.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (us *UserService) AuthenticationLocal(email, pussword string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (us *UserService) AuthenticateOAuth(token *oauth2.Token) (string, error) {
	userInfo, err := us.OUserService.GetUserInfo(token)

	if err != nil {
		return "", err
	}
	return userInfo, nil
}

func (us *UserService) CreateUserService(user *domain.User) (userdtos.UserResponse, error) {

	var userResponse userdtos.UserResponse
	err := user.Validate()

	if err != nil {
		return userResponse, err
	}

	//si existe el usuario
	foundUser, err := us.UserRepository.FindByEmail(user.Email)

	if err != nil {
		return userResponse, err
	}

	if foundUser.ID != 0 {
		return userResponse, errors.New("user already exists")
	}

	newUser, err := us.UserRepository.CreateUser(user)

	if err != nil {
		return userResponse, err
	}

	userResponse.FromEntity(*newUser)

	return userResponse, nil

}

func (us *UserService) GetUsersService() ([]userdtos.UserResponse, error) {
	users := []userdtos.UserResponse{} //arreglo para guardar rta simplificada

	usersEntity, err := us.UserRepository.GetUsers() //obtener todos los usuarios

	if err != nil {
		return users, err
	}

	//si no hay usuarios retornar error
	if len(usersEntity) == 0 {
		return users, errors.New("no users found")
	}

	//por cada usuario, crear una respuesta simplificada
	for _, user := range usersEntity {
		var userResponse userdtos.UserResponse
		userResponse.FromEntity(user)       //respuesta simplificada por cada usuario
		users = append(users, userResponse) //agregar al array de usuarios
	}

	return users, nil

}
