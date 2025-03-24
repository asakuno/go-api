package controller

import (
	"net/http"

	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	res := utils.BuildResponseSuccess("Ping successfully", gin.H{
		"message": "ok",
	})
	c.JSON(http.StatusOK, res)
}
