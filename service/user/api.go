package userservice

import (
	"context"
)

// type User interface {}

func NewUserService(repo userRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) IsUserRegistered(ctx context.Context, email string) (bool, error) {
	return s.repo.ExistsByEmail(ctx, email)
}

func (s *userService) RegisterUser(ctx context.Context, email string) error {
	return s.repo.InsertUser(ctx, email)
}
