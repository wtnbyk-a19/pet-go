package usecase

import (
	"pet-go/domain/model"
	"pet-go/domain/repository"
)

type PetUsecase interface {
	Add(*model.Pet) (err error)
	Edit(*model.Pet) (err error)
	GetPetsByUserID() (pets []*model.Pet, err error)
}

type petUsecase struct {
	petRepo repository.PetRepository
}

func NewPetUsecase(petRepo repository.PetRepository) PetUsecase {
	petUsecase := petUsecase{petRepo: petRepo}
	return &petUsecase
}

// Petの新規追加
func (usecase *petUsecase) Add(pet *model.Pet) (err error) {
	_, err = usecase.petRepo.Create(pet)
	return
}

// Petの編集
func (usecase *petUsecase) Edit(pet *model.Pet) (err error) {
	_, err = usecase.petRepo.Save(pet)
	return
}

// Petの全件検索
func (usecase *petUsecase) GetPetsByUserID() (pets []*model.Pet, err error) {
	pets, err = usecase.petRepo.FindByUserID()
	return
}
