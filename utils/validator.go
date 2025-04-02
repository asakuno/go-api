package utils

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validateOnce sync.Once
	validate     *validator.Validate
)

func GetValidator() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()

		validate.RegisterValidation("secure_password", validateSecurePassword)
	})
	return validate
}

func validateSecurePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)

	return hasUppercase && hasLowercase && hasNumber && hasSpecial
}

func ValidateRequest(validate *validator.Validate, data interface{}) (bool, string) {
	err := validate.Struct(data)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, fmt.Sprintf("予期せぬバリデーションエラーが発生しました: %v", err)
		}

		errorMessages := make([]string, 0, len(validationErrors))

		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%sは必須項目です", e.Field()))
			case "email":
				errorMessages = append(errorMessages, "有効なメールアドレスを入力してください")
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%sは%s文字以上である必要があります", e.Field(), e.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%sは%s文字以下である必要があります", e.Field(), e.Param()))
			case "secure_password":
				errorMessages = append(errorMessages, "パスワードは大文字、小文字、数字、特殊文字を含む必要があります")
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("%sの形式が正しくありません (%s)", e.Field(), e.Tag()))
			}
		}

		errorString := strings.Join(errorMessages, ", ")
		return false, errorString
	}

	return true, ""
}
