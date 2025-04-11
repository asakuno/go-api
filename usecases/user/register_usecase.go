package user_usecase

import (
	"context"
	"errors"

	"github.com/asakuno/go-api/dto/request"
	"github.com/asakuno/go-api/entities"
	"github.com/asakuno/go-api/repositories"
)

var (
	ErrUserExists = errors.New("user with the provided login ID or email already exists")
)

type (
	RegisterUsecase interface {
		Execute(ctx context.Context, req request.CreateUserRequest) (entities.User, error)
	}

	registerUsecase struct {
		userRepository repositories.UserRepository
	}
)

func NewRegisterUsecase(userRepo repositories.UserRepository) RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepo,
	}
}

func (ru *registerUsecase) Execute(ctx context.Context, req request.CreateUserRequest) (entities.User, error) {
	// Check if login ID already exists
	loginIDExists, err := ru.userRepository.CheckLoginIDExists(ctx, nil, req.LoginId)
	if err != nil {
		return entities.User{}, err
	}
	if loginIDExists {
		return entities.User{}, repositories.ErrLoginIDExists
	}

	// Check if email already exists
	emailExists, err := ru.userRepository.CheckEmailExists(ctx, nil, req.Email)
	if err != nil {
		return entities.User{}, err
	}
	if emailExists {
		return entities.User{}, repositories.ErrEmailExists
	}

	user := req.ToEntity()

	createdUser, err := ru.userRepository.Register(ctx, nil, user)
	if err != nil {
		return entities.User{}, err
	}

	return createdUser, nil
}
