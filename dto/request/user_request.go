package request

import (
	"github.com/asakuno/go-api/entities"
	"github.com/asakuno/go-api/entities/enums"
)

type (
	CreateUserRequest struct {
		LoginId  string `json:"login_id" form:"login_id" validate:"required,min=3,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (r *CreateUserRequest) ToEntity() entities.User {
	return entities.User{
		LoginId:    r.LoginId,
		Email:      r.Email,
		Password:   r.Password,
		Role:       uint8(enums.RoleUser),
		IsVerified: false,
	}
}
