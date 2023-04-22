package databases

import (
	"Backend/model"
)

// 历史创作信息的crud

// CREATE
// create path-nil History
func Create_History(user *model.MD_User) (*model.History_Create, error) {
	tmp := new(model.History_Create)
	tmp.User_Id = user.Id
	tmp.Text_path, tmp.Img_path = model.Strs{}, model.Strs{}
	user.His_Cre = *tmp
	return tmp, nil
}

// QUERY
func Query_History(user_id int) (*model.History_Create, error) {
	tmp := new(model.MD_User)
	err := model.DB.Debug().Where("Id = ?", user_id).First(&tmp).Error
	return &((*tmp).His_Cre), err
}

// UPDATE

// Insert Path: Backend Path
// all data got from Fronted have been saved as local file
func UpdateTextPath(text_path string, create *model.History_Create) (*model.History_Create, error) {
	if create != nil {
		(*create).Text_path = append((*create).Text_path, text_path)
	} else {
		(*create).Text_path = model.Strs{text_path}
	}
	return create, nil
}

func UpdateImagePath(image_path string, create *model.History_Create) (*model.History_Create, error) {
	(*create).Text_path = append((*create).Text_path, image_path)
	return create, nil
}

// DELETE
// 删除当前用户的history
func DeleteHistory(status string) error {
	return model.DB.Model(&model.MD_User{}).Where("Name = ?", status).Update("His_Cre", model.History_Create{}).Error
}
