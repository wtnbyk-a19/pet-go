package api

import (
	"pet-go/middlewares"
	"pet-go/models"
	"strconv"
	"time"

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

		adoptaversary, _ := time.Parse("2021/01/01", c.FormValue("adoptaversary"))
		birthday, _ := time.Parse("2021/01/01", c.FormValue("birthday"))

		pet := models.Pet{
			UserID:        userID,
			PetName:       c.FormValue("pet_name"),
			Gender:        c.FormValue("gender"),
			Category:      c.FormValue("category"),
			Breed:         c.FormValue("breed"),
			Adoptaversary: adoptaversary,
			Birthday:      birthday,
			Memo:          c.FormValue("memo"),
		}

		dbs.DB.Create(&pet)

		response := RegisterPetResponse{
			UserID: pet.UserID,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
