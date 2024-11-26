package main

import (
	"echo_jwt/config"
	"echo_jwt/controller"
	"echo_jwt/router"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()
	r := echo.New()
	userController := controller.NewUserController()
	router.UserRoute(r, userController)
	r.Start(":8000")
}
