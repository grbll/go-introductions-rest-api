package service

import (
	"context"
	. "github.com/grbll/go-introductions-rest-api/models"
)

type UserRepository interface {
	GetUserById(ctx context.Context, userid int) (*User, error)
	// CreateUser(ctx context.Context, user *User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) IsUserRegistered(email string) (bool, error) {
	return true, nil
}
