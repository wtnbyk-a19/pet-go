package routes

import (
	"pet-go/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.POST("/signup", api.CreateUser())
	}
}
