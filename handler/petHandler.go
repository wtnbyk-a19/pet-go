package handler

import (
	"net/http"
	"pet-go/domain/model"
	"pet-go/usecase"

	"github.com/labstack/echo"
)

type PetHandler struct {
	petUsecase usecase.PetUsecase
}

func NewPetHandler(petUsecase usecase.PetUsecase) PetHandler {
	petHandler := PetHandler{petUsecase: petUsecase}
	return petHandler
}

func (handler *PetHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pet model.Pet
		c.Bind(&pet)
		err := handler.petUsecase.Add(&pet)
		return c.JSON(http.StatusOK, err)
	}
}

func (handler *PetHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pet model.Pet
		c.Bind(&pet)
		err := handler.petUsecase.Edit(&pet)
		return c.JSON(http.StatusOK, err)
	}
}

func (handler *PetHandler) GetPetsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		pets, err := handler.petUsecase.GetPetsByUserID()
		if err != nil {
			return c.JSON(http.StatusBadRequest, pets)
		}
		return c.JSON(http.StatusOK, pets)
	}
}
