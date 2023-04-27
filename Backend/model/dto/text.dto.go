package dto

import (
	"Backend/model"
	"mime/multipart"
)

type bools []bool

type TextInput struct {
	Text string `gorm:"varchar(600)" form:"text" json:"text"`
}

type TextLabelInput struct {
	LabelList model.Strs `gorm:"longtext" form:"label_list" json:"label_list"`
}

type TextFileUpload struct {
	TextFile multipart.File `gorm:"-"`
}

type TextLabelSelect struct {
	Selection bools `gorm:"-"`
}
