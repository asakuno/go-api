package controllers

import (
	"net/http"

	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	ctx.JSON(http.StatusOK, res)
}
