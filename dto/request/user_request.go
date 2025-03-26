package request

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type (
	CreateUserRequest struct {
		LoginId  string `json:"login_id" form:"login_id" validate:"required,min=3,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
)

func (*CreateUserRequest) validateSecurePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)

	return hasUppercase && hasLowercase && hasNumber && hasSpecial
}
