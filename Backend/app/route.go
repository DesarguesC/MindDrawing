package app

// add web routes

import (
	"Backend/app/controller"
	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func AddRoutes() {
	//visit := e.Group("visit", middleware)
	e.GET("/ping", ping)
	user_api := e.Group("user")

	user_api.POST("/register", controller.Users_Register)
	user_api.POST("/login", controller.Users_Login)
	user_api.GET("/logout", controller.Users_Logout)
	user_api.GET("/account/get/all", controller.Users_GetAll)
	user_api.GET("/account/get/secq", controller.Users_GetSecA)
	user_api.POST("/account/password/pwd", controller.Users_AmendPwd_Pwd)
	user_api.POST("/account/passwprd/sec", controller.Users_AmendPwd_Sec)
}
