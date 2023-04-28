package controller

import (
	"Backend/app/response"
	"github.com/labstack/echo/v4"
)

func Text_Input_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func TextFile_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func TextLabel_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func Get_GenTexts(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}

func Editor_TextGuidance(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
}
