package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// init web service

var e *echo.Echo

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	//addRoutes()
	logrus.Info("echo framework initialized")
}
func StartServer() {
	// Backend: localhost: 3000
	e.Logger.Fatal(e.Start(":3000"))
}
