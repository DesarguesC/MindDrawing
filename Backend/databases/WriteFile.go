package databases

import (
	"bufio"
	"image"
	"os"
	"strconv"
)

var base_folder string

// 将文本写入对应用户的文件夹
func WriteText(text string, status string, num int) (string, error) {
	base_folder = "Backend/ASSETS/"
	num_str := strconv.Itoa(num)
	route := base_folder + status + "/TEXT/"
	// fmt.Println(route)
	err := os.MkdirAll(route, 0777)
	if err != nil {
		return "Text MkdirAll Failed", err
	}

	file_path := route + "story-" + num_str + ".txt"
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "OpenFile Failed", err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(text)
	writer.Flush()
	return file_path, err
}

func WriteImage(img image.Image, status string, num int) (string, error) {
	base_folder = "Backend/ASSETS/"
	num_str := strconv.Itoa(num)
	route := base_folder + status + "/IMAGE/"
	err := os.MkdirAll(route, 0777)
	if err != nil {
		return "Image MkdirAll Failed", err
	}
	img_path := route + "image-" + num_str + ".png"
	// Remain to be written...

	return img_path, err
}
