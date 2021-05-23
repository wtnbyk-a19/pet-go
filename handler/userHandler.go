package handler

import (
	"net/http"
	"pet-go/domain/model"
	"pet-go/usecase"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	userHandler := UserHandler{userUsecase: userUsecase}
	return userHandler
}

func (handler *UserHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
		err := handler.userUsecase.Add(&user)
		return c.JSON(http.StatusOK, err)
	}
}

func (handler *UserHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
		err := handler.userUsecase.Edit(&user)
		return c.JSON(http.StatusOK, err)
	}
}

func (handler *UserHandler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := handler.userUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, users)
		}
		return c.JSON(http.StatusOK, users)
	}
}
