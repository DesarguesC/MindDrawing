package model

import (
	"time"
)

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
