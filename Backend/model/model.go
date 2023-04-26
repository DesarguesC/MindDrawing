package model

import (
	"database/sql/driver"
	"strconv"
	"strings"
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
	His_Cre   History_Create `gorm:"type:varchar(500)" form:"history" json:"history"`
	LastLogin time.Time      `gorm:"type:datetime" form:"expire_time" json:"expire_time"`
	// Regist Time
}

/*
read example:

10 & Name & Email & Pwd & SecQ & SecA & 10;./TEXT/1/|./TEXT/2/;./IMAGE/1/|./IMAGE/2/ & Lastlogin

*/

func (m *MD_User) Scan(val interface{}) error {
	var err error
	s := val.([]uint8)
	ss := strings.Split(string(s), "&")
	(*m).Id, err = strconv.Atoi(ss[0])
	if err != nil {
		return err
	}
	(*m).Name = ss[1]
	(*m).Email = ss[2]
	(*m).Pwd = ss[3]
	(*m).SecQ = ss[4]
	(*m).SecA = ss[5]

	sss := strings.Split(ss[6], ";")
	(*m).His_Cre.User_Id, err = strconv.Atoi(sss[0])
	if err != nil {
		return nil
	}
	(*m).His_Cre.Text_path = strings.Split(sss[1], "|")
	(*m).His_Cre.Img_path = strings.Split(sss[2], "|")
	(*m).LastLogin, _ = time.Parse(ss[7], "2006-01-02 15:04:05")
	// string -> time

	return err
}

/*
store example:

10 & Name & Email & Pwd & SecQ & SecA & 10;./TEXT/1/|./TEXT/2/;./IMAGE/1/|./IMAGE/2/ & Lastlogin

*/

func (m MD_User) Value() (driver.Value, error) {
	str_1 := strconv.Itoa(m.Id)
	str_2 := m.Name + "&" + m.Email + "&" + m.Pwd + "&" + m.SecQ + "&" + m.SecA
	str_3 := strconv.Itoa(m.His_Cre.User_Id) + ";" + strings.Join(m.His_Cre.Text_path, "|") + ";" + strings.Join(m.His_Cre.Img_path, "|")
	str_4 := m.LastLogin.Format("2006-01-02 15:04:05")
	// time -> string

	str := str_1 + "&" + str_2 + "&" + str_3 + "&" + str_4
	return str, nil
}
