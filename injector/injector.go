package injector

import (
	"pet-go/domain/repository"
	"pet-go/handler"
	"pet-go/infrastructure/database"
	"pet-go/infrastructure/persistence"
	"pet-go/usecase"
)

func InjectDB() database.SqlHandler {
	sqlHandler := database.NewSqlHandler()
	return *sqlHandler
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	return persistence.NewUserRepository(sqlHandler)
}

func InjectUserUsecase() usecase.UserUsecase {
	UserRepo := InjectUserRepository()
	return usecase.NewUserUsecase(UserRepo)
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}
