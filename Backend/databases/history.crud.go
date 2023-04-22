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
func Query_History() {

}

// UPDATE

// Insert Path: Backend Path
// all data got from Fronted have been saved as local file
func UpdateTextPath(text_path string, create model.History_Create) (model.History_Create, error) {
	var Str model.Strs
	for i := 0; i <= len(create.Text_path); i++ {
		if i != len(create.Text_path) {
			Str[i] = create.Text_path[i]
		} else {
			Str[i] = text_path
		}

	}
	return model.History_Create{
		User_Id:   create.User_Id,
		Text_path: Str,
		Img_path:  create.Img_path,
	}, nil
}

func UpdateImagePath(image_path string, create model.History_Create) (model.History_Create, error) {
	var Str model.Strs
	for i := 0; i <= len(create.Img_path); i++ {
		if i != len(create.Img_path) {
			Str[i] = create.Text_path[i]
		} else {
			Str[i] = image_path
		}

	}
	return model.History_Create{
		User_Id:   create.User_Id,
		Text_path: create.Text_path,
		Img_path:  Str,
	}, nil
}

// DELETE
// 删除当前用户的history
func DeleteHistory(status string) error {
	return model.DB.Model(&model.MD_User{}).Where("Name = ?", status).Update("His_Cre", model.History_Create{}).Error
}
