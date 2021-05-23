package persistence

import (
	"pet-go/domain/model"
	"pet-go/domain/repository"
	"pet-go/infrastructure/database"
)

type PetRepository struct {
	database.SqlHandler
}

func NewPetRepository(sqlHandler database.SqlHandler) repository.PetRepository {
	petRepository := PetRepository{sqlHandler}
	return &petRepository
}

func (petRepo *PetRepository) Create(pet *model.Pet) (*model.Pet, error) {
	result := petRepo.SqlHandler.Conn.Create(pet)
	return pet, result.Error
}

func (petRepo *PetRepository) Save(pet *model.Pet) (*model.Pet, error) {
	result := petRepo.SqlHandler.Conn.Save(pet)
	return pet, result.Error
}

func (petRepo *PetRepository) FindByUserID() (pets []*model.Pet, err error) {
	var pet []model.Pet
	result := petRepo.SqlHandler.Conn.Find(&pet)

	for _, v := range pet {
		pets = append(pets, &v)
	}

	return pets, result.Error
}
