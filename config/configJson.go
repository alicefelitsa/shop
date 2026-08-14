package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

var Cj *viper.Viper

func init() {
	InitConfigJson("config")
}

// InitConfigJson 初始化config.json配置文件
func InitConfigJson(fileName string) {
	Cj = viper.New()
	Cj.AddConfigPath("./config/") // 文件所在目录
	Cj.SetConfigName(fileName)    // 文件名
	Cj.SetConfigType("json")      // 文件类型
	if err := Cj.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal("找不到config.json文件：", err)
		}
	}
}
