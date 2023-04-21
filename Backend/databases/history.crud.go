package databases

import (
	"Backend/model"
	"fmt"
)

// 历史创作信息的crud

// QUERY
func QueryNE(n_e string, name_mode bool) (*model.MD_User, error) {
	tmp := new(model.MD_User)
	fmt.Println(n_e)
	var err error
	if name_mode == true {
		tmp.Name = n_e
		err = model.DB.Debug().Where("Name = ?", n_e).First(&tmp).Error
	} else {
		tmp.Email = n_e
		err = model.DB.Debug().Where("Email = ?", n_e).First(&tmp).Error
	}
	return tmp, err
}

// UPDATE
// 用status记录登录的name，这里直接传入status
func UpdateUser_Text(status string, text_path string) (*model.MD_User, error) {
	var err error
	var hist model.History_Create
	var tmp model.MD_User
	fmt.Println(text_path)
	err = model.DB.Debug().Where("Name = ?", status).First(&tmp).Error
	if err != nil {
		return nil, err
	}
	hist = tmp.His_Cre
	hist, err = model.InsertTextPath(text_path, hist)
	err = model.DB.Model(&tmp).Update("His_Cre", hist).Error
	// DB.Model(&tmp)  ->  select
	return &tmp, err
}

func UpdateUser_Image(status string, image_path string) (*model.MD_User, error) {
	var err error
	var hist model.History_Create
	var tmp model.MD_User
	fmt.Println(image_path)
	err = model.DB.Debug().Where("Name = ?", status).First(&tmp).Error
	if err != nil {
		return nil, err
	}
	hist = tmp.His_Cre
	hist, err = model.InsertImagePath(image_path, hist)
	err = model.DB.Model(&tmp).Update("His_Cre", hist).Error
	return &tmp, err
}

// CREATE
// dto user直接传入，用指针返回
func CreateUser(user model.MD_User) (*model.MD_User, error) {
	err := model.DB.Debug().Create(user).Error
	return &user, err
}

// DELLETE
// 注销当前用户
func DeleteUser(status string) error {
	return model.DB.Model(&model.MD_User{}).Where("Name = ?", status).Error
}
