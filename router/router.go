package router

import (
	"echo_jwt/controller"

	"github.com/labstack/echo/v4"
)

func UserRoute(router *echo.Echo, UserController controller.UserController) {
	router.GET("/", UserController.GetUserData)
	router.POST("/user", UserController.AddUserData)
	router.PUT("/user/:id", UserController.UpdateUserData)
	router.DELETE("/user/:id", UserController.DeleteUserData)
	router.GET("/user/:id", UserController.GetUser)
	router.POST("/signin", UserController.AuthenticateUser)

}
