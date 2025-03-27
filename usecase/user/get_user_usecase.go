package user_usecase

import (
	"context"

	"github.com/asakuno/go-api/dto/response"
	"github.com/asakuno/go-api/entity/enum"
	"github.com/asakuno/go-api/repository"
)

type (
	GetUserUsecase interface {
		Execute(ctx context.Context) (response.GetAllUserResponse, error)
	}

	getUserUsecase struct {
		userRepository repository.UserRepository
	}
)

func NewGetUserUsecase(userRepo repository.UserRepository) GetUserUsecase {
	return &getUserUsecase{
		userRepository: userRepo,
	}
}

func (guu *getUserUsecase) Execute(ctx context.Context) (response.GetAllUserResponse, error) {
	users, err := guu.userRepository.GetAllUser(ctx, nil)
	if err != nil {
		return response.GetAllUserResponse{}, err
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:       user.ID,
			LoginId:  user.LoginId,
			UserRole: enum.UserRole(user.Role).GetLabel(),
		})
	}

	return response.GetAllUserResponse{
		Users: userResponses,
		Count: len(users),
	}, nil
}
