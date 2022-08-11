package config

import (
	"swordsman/logger"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("./config.json") // 指定配置文件路径
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		logger.Fatalf("Fatal error config file: %s \n", err)
	}
}
