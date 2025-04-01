package user_usecase

import (
	"context"

	"github.com/asakuno/go-api/dto/response"
	"github.com/asakuno/go-api/entities/enums"
	"github.com/asakuno/go-api/repositories"
	"github.com/google/uuid"
)

type (
	GetUserUsecase interface {
		Execute(ctx context.Context) (response.GetAllUserResponse, error)
	}

	getUserUsecase struct {
		userRepository repositories.UserRepository
	}
)

func NewGetUserUsecase(userRepo repositories.UserRepository) GetUserUsecase {
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
			ID:       uuid.UUID(user.ID),
			LoginId:  user.LoginId,
			UserRole: enums.UserRole(user.Role).GetLabel(),
		})
	}

	return response.GetAllUserResponse{
		Users: userResponses,
		Count: len(users),
	}, nil
}
