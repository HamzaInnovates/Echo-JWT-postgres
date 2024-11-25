package router

import (
	"echo_jwt/controller"

	"github.com/labstack/echo/v4"
)

func UserRoute(router *echo.Echo) {
	router.GET("/", controller.GetUserData)
	router.POST("/user", controller.AddUserData)
	router.PUT("/user/:id", controller.UpdateUserData)
	router.DELETE("/user/:id", controller.DeleteUserData)
	router.GET("/user/:id", controller.GetUser)
	router.POST("/signin", controller.AuthenticateUser)

}
