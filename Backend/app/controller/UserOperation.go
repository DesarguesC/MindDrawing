package controller

import (
	"Backend/app/middleware/validation"
	"Backend/app/response"
	"Backend/databases"
	"Backend/model"
	"Backend/model/dto"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
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
	data := new(dto.RegisterInput)
	if err := c.Bind(data); err != nil {
		return response.SendResponse(c, 400, "Bind Fained")
	}
	if (*data).Name == "" {
		return response.SendResponse(c, -200, "用户名字段不能为空")
	}
	if (*data).Email == "" {
		return response.SendResponse(c, -201, "邮箱字段不能为空")
	}
	if (*data).Pwd == "" {
		return response.SendResponse(c, -202, "密码不能为空")
	}
	if (*data).Pwd_confirm == "" {
		return response.SendResponse(c, -203, "请确认密码")
	}
	if (*data).SecQ == "" {
		return response.SendResponse(c, -204, "请输入密保问题")
	}
	if (*data).SecQ == "" {
		return response.SendResponse(c, -205, "请输入密保答案")
	}
	if (*data).Pwd != (*data).Pwd_confirm {
		return response.SendResponse(c, -102, "两次输入密码不一致")
	}
	var err error
	tmp := validation.CustomValidator{}
	t1 := validation.RegisterNameValid{Name: (*data).Name}
	t2 := validation.RegisterEmailValid{Email: (*data).Email}
	if err = tmp.Validate(t1); err != nil {
		return response.SendResponse(c, -100, "Name校验失败", t1.Name)
		// 这里使用response为前端提供信息
	}
	if err = tmp.Validate(t2); err != nil {
		return response.SendResponse(c, -101, "Email校验失败", t2.Email)
	}
	if _, err = databases.QueryNE(t1.Name, true); err == nil {
		return response.SendResponse(c, -103, "用户名已被使用")
	}
	if _, err = databases.QueryNE(t2.Email, false); err == nil {
		return response.SendResponse(c, -103, "邮箱已被使用")
	}

	user := new(model.MD_User)
	user.Name = (*data).Name
	user.Email = (*data).Email
	user.Pwd = (*data).Pwd
	user.SecQ = (*data).SecQ
	user.SecA = (*data).SecA
	user.His_Cre = model.History_Create{}
	user.LastLogin = time.Now()
	_, err = databases.CreateUser(user)
	if err == nil {
		return response.SendResponse(c, 400, "Create Failed")
	}
	return response.SendResponse(c, 000, "注册成功，请返回登录", user.Name, user.Email, user.Id)
}

func Users_Login(c echo.Context) error {
	if status != "nil" {
		return response.SendResponse(c, -333, "请勿重复登录")
	}
	data1 := new(validation.LoginValid_E)
	data2 := new(validation.LoginValid_N)
	if err := c.Bind(data1); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	data2.Name = data1.Email
	data2.Pwd = data1.Pwd
	tmp := validation.CustomValidator{}

	err1, err2 := tmp.Validate(data1), tmp.Validate(data2)
	if err1 != nil && err2 != nil {
		return response.SendResponse(c, -102, "Valid Failed")
	}

	tmp_1, err1 := databases.QueryNE(data1.Email, true)
	tmp_2, err2 := databases.QueryNE(data2.Name, false)

	if err1 == nil && err2 == nil {
		return response.SendResponse(c, -111, "不存在的用户")
	}

	// 验证密码
	if err1 != nil && tmp_1.Pwd != data1.Pwd || err2 != nil && tmp_2.Pwd != data2.Pwd {
		return response.SendResponse(c, -300, "密码错误")
	} else {
		if err1 != nil {
			status = tmp_1.Name
		} else {
			status = tmp_2.Name
		}
		return response.SendResponse(c, 001, "欢迎使用")
	}
}

func Users_Logout(c echo.Context) error {
	status = "nil"
	return response.SendResponse(c, 004, "您已退出账户")
}

func Users_GetAll(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	tmp, err := databases.QueryNE(status, true)
	if err != nil {
		logrus.Fatal("Unkown Error")
	}
	return response.SendResponse(c, 999, "用户信息", tmp.Name, tmp.Email, tmp.Id)
}

func Users_GetSecA(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, -999, "请先登录")
	}
	tmp, err := databases.QueryNE(status, true)
	if err != nil {
		logrus.Fatal("Unkown Error")
	}
	return response.SendResponse(c, 998, "密保问题", tmp.SecA)
}

func Users_AmendPwd_Pwd(c echo.Context) error {
	data := new(dto.OriPwdInput)
	if err := c.Bind(data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	data_n := validation.OriPwdValid_N{data.N_E, data.Pwd_ori, data.Pwd_new}
	data_e := validation.OriPwdValid_E{data.N_E, data.Pwd_ori, data.Pwd_new}

	tmp := validation.CustomValidator{}
	err_n, err_e := tmp.Validate(data_n), tmp.Validate(data_e)

	if err_n != nil && err_e != nil {
		return response.SendResponse(c, -104, "用户不存在")
	}

	tmp_n, err_n := databases.QueryNE(data.N_E, true)
	tmp_e, err_e := databases.QueryNE(data.N_E, false)

	if err_n != nil && tmp_n.Pwd != data_n.Pwd_ori || err_e != nil && tmp_e.Pwd != data_e.Pwd_ori {
		return response.SendResponse(c, -105, "原始密码错误")
	} else {
		if err_n != nil {
			status = tmp_n.Name
			databases.UpdateUser_Pwd(status, data.Pwd_new)
		} else {
			status = tmp_e.Name
			databases.UpdateUser_Pwd(status, data.Pwd_new)
		}
		return response.SendResponse(c, 002, "通过输入原始密码成功修改密码")
	}

}

func Users_AmendPwd_Sec(c echo.Context) error {
	data := new(dto.SecPwdInput)
	if err := c.Bind(data); err != nil {
		return response.SendResponse(c, 400, "Bind Failed")
	}
	data_n := new(validation.SecPwdValid_N)
	data_e := new(validation.SecPwdValid_E)

	tmp := validation.CustomValidator{}
	err_n, err_e := tmp.Validate(data_n), tmp.Validate(data_e)

	if err_n != nil && err_e != nil {
		return response.SendResponse(c, -104, "用户不存在")
	}

	tmp_n, err_n := databases.QueryNE(data.N_E, true)
	tmp_e, err_e := databases.QueryNE(data.N_E, false)

	if err_n != nil && tmp_n.SecA != data_n.SecA || err_e != nil && tmp_e.SecA != data_e.SecA {
		return response.SendResponse(c, -106, "密保问题错误", data.SecA)
	} else {
		if err_n != nil {
			status = tmp_n.Name
			databases.UpdateUser_Pwd(status, data.Pwd_new)
		} else {
			status = tmp_e.Name
			databases.UpdateUser_Pwd(status, data.Pwd_new)
		}
		return response.SendResponse(c, 003, "通过密保问题成功修改密码")
	}

}
