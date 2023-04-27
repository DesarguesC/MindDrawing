package config

import (
	"Backend/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

const LENGTH = 10

func Get_labels() (model.Strs, error) {
	viper.SetConfigName("labels")
	viper.SetConfigType("yaml") // yaml also be fine
	viper.AddConfigPath("./model/config/")
	// 数据库用户登录的配置文件直接放在short-link-backend文件夹内
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("Error occurred when fetching text labels")
		return model.Strs{}, err
	}
	Label := viper.GetStringMapString("labels")
	labels := new(model.Strs)
	for i := 1; i <= LENGTH; i++ {
		(*labels) = append((*labels), Label["choice_"+strconv.Itoa(i)])
	}
	return *labels, nil
}
