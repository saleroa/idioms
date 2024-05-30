package boot

import (
	"questionplatform/global"

	"github.com/spf13/viper"
)

func InitViper(ConfigPath string) {
	// 设置配置文件的名字（不需要扩展名）
	viper.SetConfigFile(ConfigPath)
	// 设置配置文件的格式
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
