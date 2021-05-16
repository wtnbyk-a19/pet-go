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

		user := models.User{
			UUID:     createUUID().String(),
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		dbs.DB.Create(&user)

		return c.JSON(fasthttp.StatusOK, nil)
	}
}
