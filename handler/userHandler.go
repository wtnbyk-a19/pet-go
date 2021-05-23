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

		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")

		user.Name = name
		user.Email = email
		user.Password = password

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
