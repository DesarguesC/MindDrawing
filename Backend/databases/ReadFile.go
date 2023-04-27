package databases

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"image"
	"io/ioutil"
	"os"
	"strconv"
)

// 从用户文件夹中读取路径文本信息，原样返回其中文本
func ReadTexts(status string, story_num int, story_name string) ([]string, error) {
	var texts []string
	base_folder = "Backend/ASSETS/"
	story_num_str := strconv.Itoa(story_num)

	// store text in History
	// ASSETS/cd/TEXT/story-1-#deep sea#/para-1.txt
	// ASSETS/cd/TEXT/story-1-#deep sea#/para-2.txt
	//	...
	//	space -> '%'  ???

	folder_route := base_folder + status + "/TEXT/story-" + story_num_str + "-#" + story_name + "#/"
	content, err := ioutil.ReadDir(folder_route)
	if err != nil {
		return []string{"ReadFolder Failed"}, err
	}
	for _, file := range content {
		reader, err := os.OpenFile(folder_route+file.Name(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			logrus.Fatal(err)
			return []string{}, err
		}
		defer reader.Close()
		scanner := bufio.NewScanner(reader)
		tmp := ""
		for scanner.Scan() {
			tmp += scanner.Text()
		}
		texts = append(texts, tmp)
	}

	fmt.Println(texts)
	return texts, err
}

// 读取用户某个故事对应的所有图片
func ReadImages(status string, story_num int, story_name string) ([]image.Image, error) {
	var image_array []image.Image
	image_array = nil

	// store images in History
	// ASSETS/cd/IMAGE/story-1-#deep sea#/image-1.txt
	// ASSETS/cd/IMAGE/story-1-#deep sea#/image-2.txt
	//	...
	//	space -> '%'  ???.

	base_folder = "Backend/Assets/"
	story_num_str := strconv.Itoa(story_num)
	//para_num_str := strconv.Itoa(para_num)

	folder_route := base_folder + status + "/IMAGE/story-" + story_num_str + "-#" + story_name + "#/"
	FileList, err := ioutil.ReadDir(folder_route)
	if err != nil {
		return nil, err
	}
	for _, file := range FileList {
		reader, err := os.Open(folder_route + file.Name())
		if err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		defer reader.Close()
		img, _, err := image.Decode(reader)
		if err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		image_array = append(image_array, img)
	}
	return image_array, err
}
