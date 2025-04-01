package repositories

import (
	"context"

	"github.com/asakuno/go-api/entities"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetAllUser(ctx context.Context, tx *gorm.DB) ([]entities.User, error)
		Register(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
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

func (ur *userRepository) Register(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error) {
	if tx == nil {
		tx = ur.db
	}

	result := tx.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil

}

func (ur *userRepository) GetAllUser(ctx context.Context, tx *gorm.DB) ([]entities.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var users []entities.User

	query := tx.WithContext(ctx).Model(&entities.User{})
	if err := query.Order("role").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
