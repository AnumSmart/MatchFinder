package auth

import "fmt"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) SayHello() string {
	return fmt.Sprintf("hello from service layer: %s", "HO-HO-HO")
}
