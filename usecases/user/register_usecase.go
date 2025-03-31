package user_usecase

import (
	"context"

	"github.com/asakuno/go-api/dto/request"
	"github.com/asakuno/go-api/entities"
	"github.com/asakuno/go-api/repositories"
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
	user := req.ToEntity()

	return user, nil
}
