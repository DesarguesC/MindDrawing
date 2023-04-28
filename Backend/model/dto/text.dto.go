package dto

import (
	"Backend/model"
	"Backend/model/config"
	"mime/multipart"
)

type bools []bool

type TextInput struct {
	Text string `gorm:"varchar(600)" form:"text" json:"text"`
}

type TextLabelIndexInput struct {
	LableList []int `gorm:"ints" form:"label_index" json:"label_index"`
}

type TextLabelInput struct {
	LabelList model.Strs `gorm:"longtext" form:"label_list" json:"label_list"`
}

type TextFileUpload struct {
	TextFile multipart.File `gorm:"-"`
}

func (m TextLabelIndexInput) GetContex() (*TextLabelInput, error) {
	Label := new(TextLabelInput)
	var err error
	(*Label).LabelList, err = config.Get_Contex(m.LableList)
	return Label, err
}

func TextTrans(labels []string) string {
	result := ""
	am := len(labels)
	for i := 0; i < am; i++ {

		// Core Function Remained to be Define!

		result += labels[am]
		if i != am-1 {
			result += ","
		}

	}
	return result
}
