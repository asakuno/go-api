package provider

import (
	"github.com/asakuno/go-api/constants"
	"github.com/asakuno/go-api/controller"
	"github.com/asakuno/go-api/repository"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	// repository層
	userRepository := repository.NewUserRepository(db)

	// controller層
	do.Provide(injector, func(i *do.Injector) (controller.UserController, error) {
		return controller.NewUserController(userRepository), nil
	})
}
