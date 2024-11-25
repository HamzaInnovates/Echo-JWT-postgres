package main

import (
	"echo_jwt/config"
	"echo_jwt/router"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()
	r := echo.New()
	router.UserRoute(r)
	r.Start(":8000")
}
