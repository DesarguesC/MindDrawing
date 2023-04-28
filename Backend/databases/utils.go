package databases

import (
	"bufio"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"path"
	"strconv"
	"strings"
)

func S(x int64) string {

	if x < 8 {
		return strconv.Itoa(int(x)) + "byte"
	} else if x < 8*1024 {
		return strconv.Itoa(int(x)/8) + "B"
	} else {
		return strconv.Itoa(int(x)/(8*1024)) + "KB"
	}
}

func File2Image(header *multipart.FileHeader) (image image.Image, err error) {
	file, err := header.Open()
	ext := strings.ToLower(path.Ext(header.Filename))
	switch ext {
	case "jpeg", "jpg":
		image, err = jpeg.Decode(bufio.NewReader(file))
	case "png":
		image, err = png.Decode(bufio.NewReader(file))
	}
	return image, err
}
