package request

import "sotoon/internal/core/dto"

type CreateUserRequest struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}

func (r *CreateUserRequest) MapToDTO() *dto.CreateUserRequest {
	return &dto.CreateUserRequest{
		Name:      r.Name,
		Cellphone: r.Cellphone,
	}
}
