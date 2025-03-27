package provider

import (
	"github.com/asakuno/go-api/constants"
	"github.com/asakuno/go-api/controller"
	"github.com/asakuno/go-api/repository"
	user_usecase "github.com/asakuno/go-api/usecase/user"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	// repository層
	userRepository := repository.NewUserRepository(db)

	// usecase層
	getUserUsecase := user_usecase.NewGetUserUsecase(userRepository)

	// controller層
	do.Provide(injector, func(i *do.Injector) (controller.UserController, error) {
		return controller.NewUserController(getUserUsecase), nil
	})
}
