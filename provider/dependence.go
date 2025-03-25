package provider

import (
	"github.com/asakuno/go-api/controller"
	"github.com/samber/do"
)

func ProvideDependencies(injector *do.Injector) {
	//controller層
	do.Provide(injector, func(i *do.Injector) (controller.UserController, error) {
		return controller.NewUserController(), nil
	})
}
