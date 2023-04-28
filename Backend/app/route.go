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
	AddUserRoutes()
	AddPicRoutes()
	AddEditorRoutes()
}

func AddUserRoutes() {
	user_api := e.Group("user")

	user_api.POST("/register", controller.Users_Register)
	user_api.POST("/login", controller.Users_Login)
	user_api.GET("/logout", controller.Users_Logout)
	user_api.GET("/account/get/all", controller.Users_GetAll)
	user_api.GET("/account/get/secq", controller.Users_GetSecA)
	user_api.POST("/account/password/pwd", controller.Users_AmendPwd_Pwd)
	user_api.POST("/account/password/sec", controller.Users_AmendPwd_Sec)
}

func AddPicRoutes() {
	gen_api := e.Group("picbook")

	gen_api.GET("/history/title", controller.Pic_Get_TitleList)
	gen_api.POST("/history/story/pic", controller.Pic_Get_ImageList)
	gen_api.POST("/history/story/text", controller.Pic_Get_TextList)

	gen_api.POST("/text/input_text", controller.Text_Input_Upload)
	gen_api.POST("/text/text_file", controller.TextFile_Upload)
	gen_api.GET("/text/get_label", controller.TextLabel_Upload)

	gen_api.POST("/image/style_img/upload", controller.StyleImage_Upload)
	gen_api.GET("/get/gen_images", controller.GenImages)
	gen_api.GET("/get/gen_texts", controller.GenTexts)

}

func AddEditorRoutes() {
	editor_api := e.Group("editor")

	editor_api.POST("/image/sketch_img/upload", controller.SketchImage_Upload)
	editor_api.POST("/text/text_input", controller.TextGuidanceEditor)
	editor_api.GET("/edit", controller.EditImage)

}
