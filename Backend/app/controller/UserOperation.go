package controller

import (
	"Backend/app/response"
	"github.com/labstack/echo/v4"
)

var status string

// 从前端bind数据到user.dto结构，调用该部分的crud
func Users_Judge(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	return response.SendResponse(c, -900, "登录状态正常")
}

func Users_Register(c echo.Context) error {
	if status != "nil" {
		return Users_Judge(c)
	}
	data := new(model.)

}
