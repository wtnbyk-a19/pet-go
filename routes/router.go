package routes

import (
	"pet-go/injector"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	userHandler := injector.InjectUserHandler()

	g := e.Group("/api")
	{
		g.GET("/users", userHandler.View())
		g.POST("/signup", userHandler.Add())
	}
}
