package usecase

import (
	"pet-go/domain/model"
	"pet-go/domain/repository"
)

type UserUsecase interface {
	Add(*model.User) (err error)
	Edit(*model.User) (err error)
	View() (users []*model.User, err error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: userRepo}
	return &userUsecase
}

// Userの新規追加
func (usecase *userUsecase) Add(user *model.User) (err error) {
	user.GenerateUUID()
	_, err = usecase.userRepo.Create(user)
	return
}

// Userの編集
func (usecase *userUsecase) Edit(user *model.User) (err error) {
	_, err = usecase.userRepo.Save(user)
	return
}

func (usecase *userUsecase) View() (users []*model.User, err error) {
	users, err = usecase.userRepo.FindAll()
	return
}
