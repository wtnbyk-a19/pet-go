package repository

import "pet-go/domain/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
}
