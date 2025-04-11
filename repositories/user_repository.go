package repositories

import (
	"context"
	"errors"

	"github.com/asakuno/go-api/entities"
	"gorm.io/gorm"
)

var (
	ErrLoginIDExists = errors.New("login_id already exists")
	ErrEmailExists   = errors.New("email already exists")
)

type (
	UserRepository interface {
		GetAllUser(ctx context.Context, tx *gorm.DB) ([]entities.User, error)
		Register(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
		CheckLoginIDExists(ctx context.Context, tx *gorm.DB, loginID string) (bool, error)
		CheckEmailExists(ctx context.Context, tx *gorm.DB, email string) (bool, error)
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

func (ur *userRepository) CheckLoginIDExists(ctx context.Context, tx *gorm.DB, loginID string) (bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var count int64
	err := tx.WithContext(ctx).Model(&entities.User{}).Where("login_id = ?", loginID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (ur *userRepository) CheckEmailExists(ctx context.Context, tx *gorm.DB, email string) (bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var count int64
	err := tx.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
