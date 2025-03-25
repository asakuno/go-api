package repository

import (
	"context"

	"github.com/asakuno/go-api/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetAllUser(ctx context.Context, tx *gorm.DB) ([]entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetAllUser(ctx context.Context, tx *gorm.DB) ([]entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var users []entity.User

	query := tx.WithContext(ctx).Model(&entity.User{})
	if err := query.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
