package controller

import (
	"Backend/app/response"
	"Backend/databases"
	"Backend/model/dto"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Pic_Get_TitleList(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	user, err := databases.QueryNE(Status, true)
	if err != nil {
		logrus.Fatal(err)
	}
	titles := (*user).His_Cre.Names
	return response.SendResponse(c, 510, "历史创作故事目录", titles)
}

func Pic_Get_ImageList(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	data := new(dto.TitleInput)
	if err := c.Bind(&data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	user, err := databases.QueryNE(Status, true)
	if err != nil {
		logrus.Fatal(err)
	}
	image_path := user.His_Cre.Img_path
	names := user.His_Cre.Names
	if len(image_path) != len(names) {
		logrus.Fatal("Serve Error")
	}
	amount := len(names)
	for i := 0; i < amount; i++ {
		if names[i] == data.Title {
			images, err := databases.ReadImages(Status, i+1, data.Title)
			if err != nil {
				logrus.Fatal("Fatal Error occurred when reading images")
			}
			return response.SendResponse(c, 555, "已获得标题为\""+data.Title+"\"的绘本插图", images)
		}
	}
	return response.SendResponse(c, -555, "当前用户不存在标题为\""+data.Title+"\"的绘本故事", Status, data.Title)
	
}

func Pic_Get_TextList(c echo.Context) error {
	if Status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	data := new(dto.TitleInput)
	if err := c.Bind(&data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	user, err := databases.QueryNE(Status, true)
	if err != nil {
		logrus.Fatal(err)
	}
	text_path := user.His_Cre.Text_path
	names := user.His_Cre.Names
	if len(names) != len(text_path) {
		logrus.Fatal("Serve Error")
	}
	amount := len(names)
	for i := 0; i < amount; i++ {
		if names[i] == data.Title {
			texts, err := databases.ReadTexts(Status, i+1, data.Title)
			if err != nil {
				logrus.Fatal("Fatal Error occurred when reading files")
			}
			return response.SendResponse(c, 556, "标题为\""+data.Title+"\"的故事对应文段如下", texts)
		}
	}
	return response.SendResponse(c, -556, "当前用户未找到标题为\""+data.Title+"\"的对应文段", Status, data.Title)
}
