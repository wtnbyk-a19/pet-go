package routes

import (
	"pet-go/injector"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	userHandler := injector.InjectUserHandler()
	petHandler := injector.InjectPetHandler()

	g := e.Group("/api")
	{
		g.GET("/user", userHandler.GetUser())
		g.GET("/users", userHandler.GetUsers())
		g.POST("/signup", userHandler.Add())
		g.POST("/user/pet", petHandler.Add())
	}
}
