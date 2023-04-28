package dto

import (
	"image"
	"mime/multipart"
)

type StyleImageUpload_Input struct {
	StyleImage *multipart.FileHeader `gorm:"-" form:"style" json:"style"`
}

type SketchImageUpload_Input struct {
	SketchImage multipart.FileHeader `gorm:"-" form:"sketch" json:"json"`
}

type StyleImageUpload struct {
	StyleImage *image.Image `gorm:"-" form:"style" json:"style"`
}

type SketchImageUpload struct {
	SketchImage *image.Image `gorm:"-" form:"sketch" json:"sketch"`
}
