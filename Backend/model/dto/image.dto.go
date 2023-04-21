package dto

import "image"

type StyleImageUpload struct {
	StyleImage image.Image `gorm:"-"`
}

type SketchImageUpload struct {
	SketchImage image.Image `gorm:"-"`
}
