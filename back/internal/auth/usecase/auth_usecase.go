package usecase

import (
	"go_real_time_chat/internal/auth/domain"
	"go_real_time_chat/internal/auth/infrastructure"

	"golang.org/x/oauth2"
)

type AuthUseCase struct {
	UserRepository infrastructure.UserRepository
}

func NewAuthUseCase(userRepo infrastructure.UserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: userRepo}
}

func (uc *AuthUseCase) AuthenticationLocal(email, password string) (*domain.User, error)

func (uc *AuthUseCase) AuthenticateOAuth(token *oauth2.Token)
