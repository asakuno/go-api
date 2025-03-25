package response

import "github.com/google/uuid"

type (
	UserResponse struct {
		ID      uuid.UUID `json:"id"`
		LoginId string    `json:"login_id"`
	}

	GetAllUserResponse struct {
		Users []UserResponse `json:"users"`
		Count int            `json:"count"`
	}
)
