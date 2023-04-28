package controller

import (
	"Backend/app/response"
	"Backend/databases"
	"Backend/model/dto"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const MAX_LENGTH = 300

func Text_Input_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	data := new(dto.TextInput)
	if err := c.Bind(&data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	am := len((*data).Text)
	if am > MAX_LENGTH {
		return response.SendResponse(c, -301, "输入文本不应超过300字", am)
	}
	_, err := databases.WriteText_into_tmp(data.Text, Status)
	if err != nil {
		return response.SendResponse(c, -1024, "后端读写错误", err.Error)
	}
	return response.SendResponse(c, 300, "accepted text")
}

func TextFile_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	file, err := c.FormFile("Text_file")
	if err != nil {
		return response.SendResponse(c, 400, "Form File Read Error")
	}
	tmp_file_path := "Backend/ASSETS/%%TMP%%/" + Status + "/InText.txt"
	src, err := file.Open()
	if err != nil {
		logrus.Fatal("file open error: %s", err)
		return response.SendResponse(c, -1024, "上传文件打开错误", file.Filename)
	}
	if file.Size > 8*1024*30 {
		return response.SendResponse(c, -302, "上传文本大小不应超过30KB (文本大小: "+databases.S(file.Size)+")", file.Size)
	}
	defer src.Close()
	cp, err := os.Create(tmp_file_path)
	if err != nil {
		logrus.Fatal("TMP file create error: %s", err)
		return response.SendResponse(c, -1024, "tmp文件创建错误", tmp_file_path)
	}
	defer cp.Close()
	// copy
	if _, err := io.Copy(cp, src); err != nil {
		logrus.Fatal("Copy Error: %s", err)
		return response.SendResponse(c, -1024, "复制文件进行暂存出错")
	}
	return response.SendResponse(c, 301, "accepted text file")
}

func TextLabel_Upload(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	IN := new(dto.TextLabelIndexInput)
	if err := c.Bind(&IN); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	data, err := IN.GetContex()
	if err != nil {
		return response.SendResponse(c, -1024, "Error Occurred", err)
	}
	text := dto.TextTrans((*data).LabelList)
	/*
		Amend this part when label config confirmed
	*/

	_, err = databases.WriteText_into_tmp(text, Status)
	if err != nil {
		return response.SendResponse(c, -1024, "后端文件读写错误", err)
	}
	return response.SendResponse(c, 666, "成功获得标签", text)
}

func GenTexts(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	// write python interface first

}

func TextGuidanceEdit(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	// write python interface first
}
