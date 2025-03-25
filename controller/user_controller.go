package controller

import (
	"net/http"

	user_usecase "github.com/asakuno/go-api/usecase/user"
	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
	}

	userController struct {
		getUserUsecase user_usecase.GetUserUsecase
	}
)

func NewUserController(getUserUsecase user_usecase.GetUserUsecase) UserController {
	return &userController{
		getUserUsecase: getUserUsecase,
	}
}

func (uc *userController) Register(ctx *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	result, err := uc.getUserUsecase.Execute(ctx.Request.Context())

	if err != nil {
		res := utils.BuildResponseFailed("エラーが発生しました", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:  true,
		Message: "",
		Data:    result,
	}
	ctx.JSON(http.StatusOK, res)
}
