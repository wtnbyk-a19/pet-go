package api

import (
	"pet-go/middlewares"
	"pet-go/models"
	"strconv"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type RegisterPetResponse struct {
	UserID int `json:"user_id"`
}

func RegisterPet() echo.HandlerFunc {
	return func(c echo.Context) error {

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		userID, _ := strconv.Atoi(c.Param("id"))

		pet := models.Pet{
			UserID: userID,
		}

		dbs.DB.Create(&pet)

		response := RegisterPetResponse{
			UserID: pet.UserID,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
