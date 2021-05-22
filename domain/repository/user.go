package repository

import "pet-go/domain/models"

type UserRepository interface {
	GetUser() (*models.User, error)
}
