package userservice

import (
	"context"
)

type userRepository interface {
	// GetById(ctx context.Context, userid int) (User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	InsertUser(ctx context.Context, email string) error
}

type userService struct {
	repo userRepository
}
