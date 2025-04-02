package entities

import (
	"github.com/asakuno/go-api/entities/custom_types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID         custom_types.UUID[User] `gorm:"type:binary(16);primary_key;<-:create"`
	LoginId    string                  `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email      string                  `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password   string                  `gorm:"type:varchar(255);not null"`
	Role       uint8                   `gorm:"type:tinyint;not null;default:1"`
	IsVerified bool                    `gorm:"not null;default:false"`
	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if uuid.UUID(u.ID) == uuid.Nil {
		uid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		u.ID = custom_types.UUID[User](uid)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
