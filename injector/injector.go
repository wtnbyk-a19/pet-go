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

// User
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

// Pet
func InjectPetRepository() repository.PetRepository {
	sqlHandler := InjectDB()
	return persistence.NewPetRepository(sqlHandler)
}

func InjectPetUsecase() usecase.PetUsecase {
	PetRepo := InjectPetRepository()
	return usecase.NewPetUsecase(PetRepo)
}

func InjectPetHandler() handler.PetHandler {
	return handler.NewPetHandler(InjectPetUsecase())
}
