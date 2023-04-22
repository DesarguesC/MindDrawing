package dto

import "mime/multipart"

type bools []bool

type TextInput struct {
	Text string `gorm:"varchar(600)" form:"text" json:"text"`
}

type TextFileUpload struct {
	TextFile multipart.File `gorm:"-"`
}

type TextLabelSelect struct {
	Selection bools `gorm:"-"`
}
