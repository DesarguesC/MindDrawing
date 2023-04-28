package config

import (
	"Backend/model"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

const LENGTH = 10

func Get_labels() (model.Strs, int, error) {
	viper.SetConfigName("labels")
	viper.SetConfigType("yaml") // yaml also be fine
	viper.AddConfigPath("./model/config/")
	// 数据库用户登录的配置文件直接放在short-link-backend文件夹内
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("Error occurred when fetching text labels")
		return model.Strs{}, LENGTH, err
	}
	Label := viper.GetStringMapString("labels")
	labels := new(model.Strs)
	for i := 1; i <= LENGTH; i++ {
		(*labels) = append((*labels), Label["choice_"+strconv.Itoa(i)])
	}
	return *labels, LENGTH, nil
}

func Get_Contex(x []int) (model.Strs, error) {
	str, length, _ := Get_labels()
	re := model.Strs{}
	am := len(x)
	if am+1 < length {
		return re, errors.New("Length not Matched Error")
	}
	for i := 0; i < am; i++ {
		re = append(re, str[x[i]])
	}
	return re, nil
}
