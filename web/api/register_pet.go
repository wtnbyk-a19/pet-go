package api

import (
	"pet-go/middlewares"
	"pet-go/models"
	"time"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func RegisterPet() echo.HandlerFunc {
	return func(c echo.Context) error {

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		userUUID := c.FormValue("user_uuid")
		petName := c.FormValue("pet_name")
		gender := c.FormValue("gender")
		category := c.FormValue("category")
		breed := c.FormValue("breed")
		memo := c.FormValue("memo")
		adoptaversary, _ := time.Parse("2021/01/01", c.FormValue("adoptaversary"))
		birthday, _ := time.Parse("2021/01/01", c.FormValue("birthday"))

		pet := models.Pet{
			PetUUID:       createUUID().String(),
			UserUUID:      userUUID,
			PetName:       petName,
			Gender:        gender,
			Category:      category,
			Breed:         breed,
			Adoptaversary: adoptaversary,
			Birthday:      birthday,
			Memo:          memo,
		}

		dbs.DB.Create(&pet)

		return c.JSON(fasthttp.StatusOK, nil)
	}
}
