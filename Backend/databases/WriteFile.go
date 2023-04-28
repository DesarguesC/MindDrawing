package databases

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
)

var base_folder string

// 将输入文段写入临时文件
func WriteText_into_tmp(text string, status string) (string, error) {
	// TMP文件夹中某个用户的tmp
	file_path := "Backend/ASSETS/%%TMP%%/" + status + "/InText.txt"
	err := os.MkdirAll(file_path, 0777)
	if err != nil {
		logrus.Panic(err)
		return file_path, err
	}
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logrus.Panic(err)
		return file_path, err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(text)
	writer.Flush()

	return "", err
}

// 将单张图片写入临时文件，这里只处理上传的图片（上传只有单张图片的接口，所以这里只写入一张）
// 生成对多张图片直接通过WriteImage写入用户文件夹而不是tmp文件夹
func WriteImage_into_tmp(image image.Image, status string, usage string, format string) (string, error) {
	// usage = {"sketch", "style"}
	img_path := "Backend/ASSETS/%%TMP%%/" + status + "/image-" + usage + "." + format
	err := os.MkdirAll(img_path, 0777)
	if err != nil {
		logrus.Panic(err)
		return img_path, err
	}
	file, err := os.Create(img_path)
	if err != nil {
		logrus.Panic(err)
		return img_path, err
	}
	defer file.Close()

	if format == ".png" {
		err = png.Encode(file, nil)
	} else {
		err = jpeg.Encode(file, image, &jpeg.Options{100})
	}
	// 尽量使用jpeg格式
	return img_path, nil
}

func DeleteAll_TMP(status string) error {
	base_folder = "Backend/ASSET/%%TMP%%/status/"
	err := os.RemoveAll(base_folder)
	fmt.Println(err)
	// 每次生成结束后tmp下用户文件夹会被清空
	return err
}

// 将文本写入对应用户的文件夹
func WriteText(text string, status string, story_num int, para_num int, story_name string) (string, string, error) {
	base_folder = "Backend/ASSETS/"
	story_num_str := strconv.Itoa(story_num)
	para_num_str := strconv.Itoa(para_num)

	route := base_folder + status + "/TEXT/"

	err := os.MkdirAll(route, 0777)
	if err != nil {
		return "Text MkdirAll Failed", "nil", err
	}

	user_path := "story-" + story_num_str + "-#" + story_name + "#/para-" + para_num_str + ".txt"

	// store text in History
	// ASSETS/cd/TEXT/story-1-#deep sea#/para-1.txt
	// ASSETS/cd/TEXT/story-1-#deep sea#/para-2.txt
	//	...
	//	space -> '%'  ???

	file_path := route + user_path

	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "OpenFile Failed", "nil", err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(text)
	writer.Flush()
	return file_path, user_path, err
}

func WriteImage(img image.Image, status string, story_num int, image_num int, story_name string, format string) (string, string, error) {
	base_folder = "Backend/ASSETS/"
	story_num_str := strconv.Itoa(story_num)
	image_num_str := strconv.Itoa(image_num)
	route := base_folder + status + "/IMAGE/"
	err := os.MkdirAll(route, 0777)
	if err != nil {
		return "Image MkdirAll Failed", "nil", err
	}

	var form string
	if format == "png" || format == "PNG" {
		form = ".png"
	}
	if format == "jpeg" || format == "jpeg" {
		form = ".jpeg"
	}

	user_path := "story-" + story_num_str + "-#" + story_name + "#/image-" + image_num_str + form
	img_path := route + user_path

	// store images in History
	// ASSETS/cd/IMAGE/story-1-#deep sea#/image-1.txt
	// ASSETS/cd/IMAGE/story-1-#deep sea#/image-2.txt
	//	...
	//	space -> '%'  ???.

	file, err := os.Create(img_path)
	if err != nil {
		return img_path, "nil", err
	}
	defer file.Close()

	if form == ".png" {
		err = png.Encode(file, nil)
	} else {
		err = jpeg.Encode(file, img, &jpeg.Options{100})
	}
	return img_path, user_path, err
}
