package routes

import (
	"github.com/asakuno/go-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func User(route *gin.Engine, injector *do.Injector) {
	userController := do.MustInvoke[controllers.UserController](injector)
	routes := route.Group("/api/v1/users")

	{
		routes.GET("", userController.GetAllUser)
		routes.POST("", userController.Register)
	}
}
