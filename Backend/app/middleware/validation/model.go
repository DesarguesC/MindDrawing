package validation

type RegisterValid struct {
	Name        string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{}"`
	Email       string `form:"email" json:"email" query:"email" validate:"contains=@[.]*.com"`
	Pwd         string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_confirm string `form:"pwd_confirm" json:"pwd_confirm" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
	SecQ        string `form:"secq" json:"secq" query:"secq" validate:"excludesall=!@#$%^&*()_-{}"`
	SecA        string `form:"seca" json:"seca" query:"seca" validate:"excludesall=!@#$%^&*()_-{}"`
}

type RegisterNameValid struct {
	Name string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{}"`
}
type RegisterEmailValid struct {
	Email string `form:"email" json:"email" query:"email" validate:"contains=@[.]*.com"`
}

type LoginValid_E struct {
	Email string `form:"email" json:"email" query:"email" validate:"contains=@[.*].com"`
	Pwd   string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type LoginValid_N struct {
	Name string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd  string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type SecPwdValid_E struct {
	// 用原始密码修改
	Email   string `form:"email" json:"email" query:"email" validate:"contains=@[.*].com"`
	Pwd_ori string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_new string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type SecPwdValid_N struct {
	// 用原始密码修改
	Name    string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_ori string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_new string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type OriPwdInput_E struct {
	// 用密保修改
	Email   string `form:"email" json:"email" query:"email" validate:"contains=@[.*].com"`
	SecA    string `form:"seca" json:"seca" query:"seca" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_new string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type OriPwdInput_N struct {
	// 用密保修改
	Name    string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{}"`
	SecA    string `form:"seca" json:"seca" query:"seca" validate:"excludesall=!@#$%^&*()_-{}"`
	Pwd_new string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}
