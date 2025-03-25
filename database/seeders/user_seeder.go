package seeders

import (
	"fmt"
	"github.com/asakuno/go-api/entity"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ListUserSeeder(db *gorm.DB) error {
	gofakeit.Seed(0)
	userCount := 10

	users := make([]entity.User, userCount)

	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	for i := 0; i < userCount; i++ {
		role := uint8(1)
		if i < 5  {
			role = uint8(2)
		}

		users[i] = entity.User {
			ID:         uuid.New(),
			LoginId:    fmt.Sprintf("user%d", i+1),
			Email:      gofakeit.Email(),
			Password:   string(hashedPassword),
			Role:       role,
			IsVerified: true,
		}
	}
	
	result := db.Create(&users)
	return result.Error
}