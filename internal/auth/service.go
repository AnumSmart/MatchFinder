package auth

import (
	"fmt"
	"matchfinder/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepo,
	}
}

func (a *AuthService) SayHello() string {
	return fmt.Sprintf("hello from service layer: %s", "HO-HO-HO")
}
