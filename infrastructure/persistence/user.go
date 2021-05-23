package persistence

import (
	"pet-go/domain/model"
	"pet-go/domain/repository"
	"pet-go/infrastructure/database"
)

type UserRepository struct {
	database.SqlHandler
}

func NewUserRepository(sqlHandler database.SqlHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) Create(user *model.User) (*model.User, error) {
	result := userRepo.SqlHandler.Conn.Create(user)
	return user, result.Error
}

func (userRepo *UserRepository) Save(user *model.User) (*model.User, error) {
	result := userRepo.SqlHandler.Conn.Save(user)
	return user, result.Error
}

func (userRepo *UserRepository) FindAll() (users []*model.User, err error) {
	var user []model.User
	result := userRepo.SqlHandler.Conn.Find(&user)

	for _, v := range user {
		users = append(users, &v)
	}

	return users, result.Error
}
