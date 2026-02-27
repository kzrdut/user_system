package repository

import (
	"context"
	"user_system/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error)
	GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	DeleteUser(ctx context.Context, userID domain.UserID) error
}
