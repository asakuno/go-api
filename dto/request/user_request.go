package request

type (
	CreateUserRequest struct {
		LoginId  string `json:"login_id" form:"login_id" validate:"required,min=3,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)
