package app

// add web routes

import (
	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func AddRoutes() {
	//visit := e.Group("visit", middleware)
	api := e.Group("api")
	api.GET("/ping", ping)
}
