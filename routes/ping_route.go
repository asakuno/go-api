package routes

import (
	"github.com/asakuno/go-api/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Ping(route *gin.Engine, injector *do.Injector) {
	routes := route.Group("/pings")
	{
		routes.GET("", controller.Ping)
	}
}
