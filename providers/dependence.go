package providers

import (
	"github.com/asakuno/go-api/constants"
	"github.com/asakuno/go-api/controllers"
	"github.com/asakuno/go-api/repositories"
	"github.com/asakuno/go-api/utils"

	user_usecase "github.com/asakuno/go-api/usecases/user"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	validate := utils.GetValidator()
	// repository層
	userRepository := repositories.NewUserRepository(db)

	// usecase層
	getUserUsecase := user_usecase.NewGetUserUsecase(userRepository)
	registerUsecase := user_usecase.NewRegisterUsecase(userRepository)

	// controller層
	do.Provide(injector, func(i *do.Injector) (controllers.UserController, error) {
		return controllers.NewUserController(getUserUsecase, registerUsecase, validate), nil
	})
}
