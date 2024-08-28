package contract

import (
	"context"
	"sotoon/internal/core/dto"
	"sotoon/internal/core/entity"
)

type UserStore interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserService interface {
	Create(ctx context.Context, user *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}
