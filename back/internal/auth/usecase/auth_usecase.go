package usecase

import (
	"errors"
	"go_real_time_chat/internal/auth/domain"
	"go_real_time_chat/internal/auth/infrastructure"
	"go_real_time_chat/internal/auth/infrastructure/oauth"

	"golang.org/x/oauth2"
)

type AuthUseCase struct {
	UserRepository infrastructure.UserRepository
	OAuthService   oauth.OAuthService
}

func NewAuthUseCase(userRepo infrastructure.UserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: userRepo}
}

func (uc *AuthUseCase) AuthenticationLocal(email, password string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (uc *AuthUseCase) AuthenticateOAuth(token *oauth2.Token) (string, error) {
	userInfo, err := uc.OAuthService.GetUserInfo(token)

	if err != nil {
		return "", err
	}
	return userInfo, nil
}

func (uc *AuthUseCase) CreateUserService(user *domain.User) (*domain.User, error) {

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	//si existe el usuario
	foundUser, err := uc.UserRepository.FindByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if foundUser.ID != 0 {
		return nil, errors.New("user already exists")
	}

	return uc.UserRepository.CreateUser(user)

}
