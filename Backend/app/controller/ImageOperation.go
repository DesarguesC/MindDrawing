package controller

import (
	"Backend/app/response"
	"github.com/labstack/echo/v4"
)

func StyleImage_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}

}

func SketchImage_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func Get_GenImages(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func EditImage(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}
