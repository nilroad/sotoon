package dto

import "sotoon/internal/core/entity"

type CreateUserRequest struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}

func (r *CreateUserRequest) MapDTOToEntity() *entity.User {
	return &entity.User{
		Name:      r.Name,
		Cellphone: r.Cellphone,
	}
}

type CreateUserResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}

func MapUserEntityToCreateUserResponseDTO(user *entity.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Cellphone: user.Cellphone,
	}
}
