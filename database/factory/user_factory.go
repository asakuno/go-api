package factory

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/asakuno/go-api/entity"
)

type UserFactory struct {
	Count       int
	AdminCount  int
	DefaultRole uint8
	Password    string
}

func NewUserFactory() *UserFactory {
	return &UserFactory{
		Count:       10,
		DefaultRole: 1,
		Password:    "password",
	}
}

func (f *UserFactory) create() ([]entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	gofakeit.Seed(time.Now().UnixNano())

	users := make([]entity.User, f.Count)

	for i := 0; i < f.Count; i++ {
		role := f.DefaultRole

		users[i] = entity.User{
			ID:         uuid.New(),
			LoginId:    fmt.Sprintf("user%d", i+1),
			Email:      gofakeit.Email(),
			Password:   string(hashedPassword),
			Role:       role,
			IsVerified: true,
			Timestamp: entity.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
	}

	return users, nil
}

func (f *UserFactory) CreateAndSave(db *gorm.DB) ([]entity.User, error) {
	users, err := f.create()
	if err != nil {
		return nil, err
	}

	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			return nil, fmt.Errorf("failed to create user %d: %w", i, err)
		}
	}

	return users, nil
}
