package controllers

import (
	"net/http"

	"github.com/asakuno/go-api/dto/request"
	user_usecase "github.com/asakuno/go-api/usecases/user"
	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
	}

	userController struct {
		getUserUsecase user_usecase.GetUserUsecase
		validate       *validator.Validate
	}
)

func NewUserController(getUserUsecase user_usecase.GetUserUsecase, validate *validator.Validate) UserController {
	return &userController{
		getUserUsecase: getUserUsecase,
		validate:       validate,
	}
}

func (uc *userController) Register(ctx *gin.Context) {
	var request request.CreateUserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res := utils.BuildResponseFailed("リクエスト形式が正しくありません", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	isValid, errorMsg := utils.ValidateRequest(uc.validate, request)
	if !isValid {
		res := utils.BuildResponseFailed("入力内容に問題があります", "バリデーションエラー", errorMsg)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Usecaseあとで追加
	res := utils.BuildResponseSuccess("ユーザーが正常に登録されました", gin.H{
		"login_id": request.LoginId,
	})
	ctx.JSON(http.StatusCreated, res)
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
