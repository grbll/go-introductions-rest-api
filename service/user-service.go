package service

import (
	"context"
)

// type User interface {}

type UserRepository interface {
	// GetById(ctx context.Context, userid int) (User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	InsertUser(ctx context.Context, email string) error
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

func (s *UserService) RegisterUser(ctx context.Context, email string) error {
	return s.repo.InsertUser(ctx, email)
}
