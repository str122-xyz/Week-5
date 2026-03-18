package services

import (
	"github.com/str122-xyz/gin-firebase-backend/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{userRepo: repositories.NewUserRepository()}
}