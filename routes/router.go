package routes

import (
	"pet-go/handler"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo, userHandler handler.UserHandler) {

	g := e.Group("/api")
	{
		g.POST("/signup", userHandler.Add())
	}
}
