package injector

import (
	"pet-go/domain/repository"
	"pet-go/handler"
	"pet-go/infrastructure/database"
	"pet-go/infrastructure/persistence"
	"pet-go/usecase"
)

func InjectUserRepository() repository.UserRepository {
	dbs := database.DatabaseService()
	return persistence.NewUserRepository(dbs)
}

func InjectUserUsecase() usecase.UserUsecase {
	UserRepo := InjectUserRepository()
	return usecase.NewUserUsecase(UserRepo)
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}
