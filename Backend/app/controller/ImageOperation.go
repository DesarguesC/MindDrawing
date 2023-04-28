package controller

import (
	"Backend/app/response"
	"Backend/databases"
	"Backend/model/dto"
	"github.com/labstack/echo/v4"
)

func StyleImage_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	IN := new(dto.StyleImageUpload_Input)
	if err := c.Bind(&IN); err != nil {
		return response.SendResponse(c, 400, "Bind Faled")
	}
	if (*IN).StyleImage.Size > 20*1024*1024*8 {
		return response.SendResponse(c, -304, "图像文件大小不得超过20M", databases.S((*IN).StyleImage.Size))
	}
	img, form, err := databases.File2Image(IN.StyleImage)
	if err != nil {
		return response.SendResponse(c, -1024, "Error Occurred when reading inputs", err)
	}
	A, B := img.Bounds().Min, img.Bounds().Max
	weight, height := B.X-A.X, B.Y-A.Y

	if weight*height > 512*512 {
		return response.SendResponse(c, -305, "图像像素大小不得超过512 * 512")
	}
	_, err = databases.WriteImage_into_tmp(img, Status, "style", form)
	if err != nil {
		return response.SendResponse(c, -1024, "Error Occurred when wring image into tmp", err)
	}
	return response.SendResponse(c, 303, "style accepted")
}

func SketchImage_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	IN := new(dto.StyleImageUpload_Input)
	if err := c.Bind(&IN); err != nil {
		return response.SendResponse(c, 400, "Bind Faled")
	}
	if (*IN).StyleImage.Size > 20*1024*1024*8 {
		return response.SendResponse(c, -304, "图像文件大小不得超过20M", databases.S((*IN).StyleImage.Size))
	}
	img, form, err := databases.File2Image(IN.StyleImage)
	if err != nil {
		return response.SendResponse(c, -1024, "Error Occurred when reading inputs", err)
	}
	A, B := img.Bounds().Min, img.Bounds().Max
	weight, height := B.X-A.X, B.Y-A.Y

	if weight*height > 512*512 {
		return response.SendResponse(c, -305, "图像像素大小不得超过512 * 512")
	}
	_, err = databases.WriteImage_into_tmp(img, Status, "sketch", form)
	if err != nil {
		return response.SendResponse(c, -1024, "Error Occurred when wring image into tmp", err)
	}
	return response.SendResponse(c, 303, "sketch accepted")
}

func GenImages(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}

}

func EditImage(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}

}
