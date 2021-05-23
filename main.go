package main

import (
	"pet-go/injector"
	"pet-go/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	userHandler := injector.InjectUserHandler()

	e := echo.New()

	// Routes
	routes.Init(e, userHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
