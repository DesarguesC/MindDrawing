package model

import (
	"time"
)

type History_Create struct {
	User_Id   int  `gorm:"type:uint" form:"user_id" json:"user_id"`
	Text_path Strs `gorm:"type:longtext" form:"user_text_path" json:"user_text_path,omitempty"`
	Img_path  Strs `gorm:"type:longtext" form:"user_img_path" json:"user_img_path,omitempty"`
}

// 历史数据保存在文件中

type MD_User struct {
	Id        int            `gorm:"type:uint;primaryKey autoincrement=1000" form:"id" json:"id"`
	Name      string         `gorm:"type:varchar(20)" form:"name" json:"name"`
	Email     string         `gorm:"type:varchar(20)" form:"email" json:"email"`
	Pwd       string         `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`
	SecQ      string         `gorm:"type:varchar(100)" form:"secq" json:"secq"`
	SecA      string         `gorm:"type:varchar(100)" form:"seca" json:"seca"`
	His_Cre   History_Create `gorm:"type:history" form:"history" json:"history"`
	LastLogin time.Time      `gorm:"type:datetime" form:"expire_time" json:"expire_time"`
}
