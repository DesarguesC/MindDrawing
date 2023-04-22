package databases

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// 从用户文件夹中读取路径文本信息，原样返回其中文本
func ReadText(status string, num int) (string, error) {
	base_folder = "Backend/ASSETS/"
	num_str := strconv.Itoa(num)
	route := base_folder + status + "/TEXT/story-" + num_str + ".txt"
	content, err := ioutil.ReadFile(route)
	if err != nil {
		return "OpenFile Failed", err
	}
	if err != nil {
		return "ReadFile Failed", err
	}
	fmt.Print(string(content))
	return string(content), err
}
