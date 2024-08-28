package usersvc

import (
	"context"
	"sotoon/internal/core/contract"
	"sotoon/internal/core/dto"
)

type Service struct {
	userStore contract.UserStore
}

func New(userStore contract.UserStore) *Service {
	return &Service{
		userStore: userStore,
	}
}

func (r *Service) Create(ctx context.Context, user *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	newUser, err := r.userStore.Create(ctx, user.MapDTOToEntity())
	if err != nil {
		return nil, err
	}

	return dto.MapUserEntityToCreateUserResponseDTO(newUser), nil
}
