package routes

import (
	"github.com/asakuno/go-api/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func User(route *gin.Engine, injector *do.Injector) {
	userController := do.MustInvoke[controller.UserController](injector)
	routes := route.Group("/api/v1/user")

	{
		routes.GET("", userController.GetAllUser)
		routes.POST("", userController.Register)
	}
}
