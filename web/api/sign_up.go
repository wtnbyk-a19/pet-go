package api

import (
	"pet-go/middlewares"
	"pet-go/models"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type SignUpResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		uuid := createUUID().String()
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")

		user := models.User{
			UUID:     uuid,
			Name:     name,
			Email:    email,
			Password: password,
		}

		dbs.DB.Create(&user)

		return c.JSON(fasthttp.StatusOK, nil)
	}
}
