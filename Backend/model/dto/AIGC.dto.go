package dto

type TitleInput struct {
	Title string `gorm:"type:varchar(50)" form:"title_name" json:"title_name"`
}
