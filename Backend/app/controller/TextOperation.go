package controller

import (
	"Backend/app/response"
	"Backend/model/dto"
	"github.com/labstack/echo/v4"
)

func Text_Input_Upload(c echo.Context) error {
	// 直接输入文本
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	data := new(dto.TextInput)
	if err := c.Bind(&data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
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
