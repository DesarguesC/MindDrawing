package dto

type SecPwdInput struct {
	// 用输入密保答案修改
	N_E     string `gorm:"type:varchar(20)" form:"n_e" json:"n_e"`
	SecA    string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
	Pwd_new string `gorm:"type:varchar(100)" form:"pwd_new" json:"pwd_new"`
}

type OriPwdInput struct {
	// 用输入原始密码修改
	N_E     string `gorm:"type:varchar(20)" form:"n_e" json:"n_e"`
	Pwd_ori string `gorm:"type:varchar(100)" form:"pwd_ori" json:"pwd_ori"`
	Pwd_new string `gorm:"type:varchar(100)" form:"pwd_new" json:"pwd_new"`
}

type LoginInput struct {
	N_E string `gorm:"type:varchar(20)" form:"n_e" json:"n_e"`
	Pwd string `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`
}

type RegisterInput struct {
	//Id          int    `gorm:"type:uint;primaryKey autoincrement=1000" form:"id" json:"id"` // 自动生成
	Name        string `gorm:"type:varchar(20)" form:"name" json:"name"`
	Email       string `gorm:"type:varchar(20)" form:"email" json:"email"`
	Pwd         string `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`
	Pwd_confirm string `gorm:"type:varchar(100)" form:"pwd_confirm" json:"pwd_confirm"`
	SecQ        string `gorm:"type:varchar(100)" form:"secq" json:"secq"`
	SecA        string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
}
