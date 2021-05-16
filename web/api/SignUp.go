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

		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")

		user := models.User{
			UUID:     createUUID().String(),
			Name:     name,
			Email:    email,
			Password: password,
		}

		dbs.DB.Create(&user)

		response := SignUpResponse{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
