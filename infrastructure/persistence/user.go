package persistence

import (
	"pet-go/domain/model"
	"pet-go/domain/repository"
	"pet-go/infrastructure/database"
)

type UserRepository struct {
	dbs *database.DatabaseClient
}

func NewUserRepository(dbs *database.DatabaseClient) repository.UserRepository {
	userRepository := UserRepository{dbs}
	return &userRepository
}

func (userRepo *UserRepository) Create(user *model.User) (*model.User, error) {
	result := userRepo.dbs.DB.Create(user)
	defer userRepo.dbs.DB.Close()
	return user, result.Error
}

func (userRepo *UserRepository) Save(user *model.User) (*model.User, error) {
	result := userRepo.dbs.DB.Save(user)
	defer userRepo.dbs.DB.Close()
	return user, result.Error
}

func (userRepo *UserRepository) Update(user *model.User, mp map[string]interface{}) (*model.User, error) {
	result := userRepo.dbs.DB.Model(user).Updates(mp)
	defer userRepo.dbs.DB.Close()
	return user, result.Error
}