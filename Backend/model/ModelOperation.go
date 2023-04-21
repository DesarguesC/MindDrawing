package model

// Insert Path: Backend Path
// all data got from Fronted have been saved as local file

func InsertTextPath(text_path string, create History_Create) (History_Create, error) {
	var Str Strs
	for i := 0; i <= len(create.Text_path); i++ {
		if i != len(create.Text_path) {
			Str[i] = create.Text_path[i]
		} else {
			Str[i] = text_path
		}

	}
	return History_Create{
		User_Id:   create.User_Id,
		Text_path: Str,
		Img_path:  create.Img_path,
	}, nil
}

func InsertImagePath(image_path string, create History_Create) (History_Create, error) {
	var Str Strs
	for i := 0; i <= len(create.Img_path); i++ {
		if i != len(create.Img_path) {
			Str[i] = create.Text_path[i]
		} else {
			Str[i] = image_path
		}

	}
	return History_Create{
		User_Id:   create.User_Id,
		Text_path: create.Text_path,
		Img_path:  Str,
	}, nil
}
