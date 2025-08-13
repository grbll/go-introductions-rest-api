package service

import (
	"context"
	. "github.com/grbll/go-introductions-rest-api/models"
)

type UserRepository interface {
	GetById(ctx context.Context, userid int) (*User, error)
	// CreateUser(ctx context.Context, user *User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) IsUserRegistered(ctx context.Context, email string) (bool, error) {
	return s.repo.ExistsByEmail(ctx, email)
}
