package controller

import (
	"net/http"

	"github.com/asakuno/go-api/dto/response"
	"github.com/asakuno/go-api/repository"
	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
	}

	userController struct {
		userRepository repository.UserRepository
	}
)

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{
		userRepository: userRepo,
	}
}

func (uc *userController) Register(ctx *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	users, err := uc.userRepository.GetAllUser(ctx.Request.Context(), nil)

	if err != nil {
		res := utils.BuildResponseFailed("エラーが発生しました", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:      user.ID,
			LoginId: user.LoginId,
		})
	}

	resData := response.GetAllUserResponse{
		Users: userResponses,
		Count: len(users),
	}

	res := utils.Response{
		Status:  true,
		Message: "",
		Data:    resData,
	}
	ctx.JSON(http.StatusOK, res)
}
