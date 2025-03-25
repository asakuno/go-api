package controller

import (
	"net/http"

	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
	}

	userController struct {
	}
)

func NewUserController() UserController {
	return &userController{}
}

func (uc *userController) Register(ctx *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	ctx.JSON(http.StatusOK, res)
}
