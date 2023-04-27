package model

import (
	"database/sql/driver"
	"strconv"
	"strings"
)

type History_Create struct {
	User_Id   int  `gorm:"type:uint" form:"user_id" json:"user_id"`
	Text_path Strs `gorm:"type:longtext" form:"user_text_path" json:"user_text_path,omitempty"`
	Img_path  Strs `gorm:"type:longtext" form:"user_img_path" json:"user_img_path,omitempty"`
	Names     Strs `gorm:"type:longtext" form:"user_story_names" json:"user_story_names,omitempty"`
}

/*
History_Create Storage:

for example:
one = History_Create{User_Id: 10, Text_path: Strs{'./TEXT/1/', './TEXT/2/'}, Img_path: Strs{'./IMAGE/1/', './IMAGE/2/'}}

read in database as:
10;./TEXT/1/|./TEXT/2/;./IMAGE/1/|./IMAGE/2/


*/

func (m *History_Create) Scan(val interface{}) error {
	var err error
	s := val.([]uint8)
	ss := strings.Split(string(s), ";")
	(*m).User_Id, err = strconv.Atoi(ss[0])
	(*m).Text_path = strings.Split(ss[1], "|")
	(*m).Img_path = strings.Split(ss[2], "|")
	(*m).Names = strings.Split(ss[3], "|")
	return err
}

func (m History_Create) Value() (driver.Value, error) {
	str_1 := strconv.Itoa(m.User_Id)
	str_2 := strings.Join(m.Text_path, "|")
	str_3 := strings.Join(m.Img_path, "|")
	str_4 := strings.Join(m.Names, "|")
	str := str_1 + ";" + str_2 + ";" + str_3 + ";" + str_4
	return str, nil
}
