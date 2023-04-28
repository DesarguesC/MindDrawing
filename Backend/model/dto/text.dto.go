package dto

import (
	"mime/multipart"
)

type bools []bool
type ints []int

type TextInput struct {
	Text string `gorm:"varchar(600)" form:"text" json:"text"`
}

type TextLabelInput struct {
	LabelList ints `gorm:"longtext" form:"label_list" json:"label_list"`
	// index list
}

type TextFileUpload struct {
	TextFile *multipart.File `gorm:"-" form:"text_input" json:"text_input"`
}

type TextLabelSelect struct {
	Selection bools `gorm:"-"`
}
