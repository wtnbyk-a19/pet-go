package repository

import "pet-go/domain/model"

type PetRepository interface {
	Create(pet *model.Pet) (*model.Pet, error)
	Save(pet *model.Pet) (*model.Pet, error)
	FindByUserID() (pets []*model.Pet, err error)
}
